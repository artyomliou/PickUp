#!/bin/sh

filled() {
    file=$1
    key=$2
    val="$(awk -F "=" "/$key/ {print \$2}" $file)"
    if [ -n "$val" ]; then
        return 0
    else
        return 1
    fi
}
get_val() {
    file=$1
    key=$2
    awk -F "=" "/$key/ {print \$2}" $file
}
fill_val() {
    file=$1
    key=$2
    val=$3
    if is_key_set $file $key; then
        sed -i -E "s/($key=).*/\1\"$val\"/" $file
    else
        echo "$key=$val" | tee -a $file >/dev/null
    fi
}
is_key_set() {
    file=$1
    key=$2
    grep "$key" $file >/dev/null
}
gen_pwd() {
    head /dev/urandom | tr -dc A-Za-z0-9 | head -c10
}
gen_app_key() {
    openssl rand -base64 32
}

##################
# .env.docker
##################
DOCKER=".env.docker"
if [ ! -e $DOCKER ]; then
    cp .env.example $DOCKER
fi

# Generate secret variable for local use
if ! filled $DOCKER MYSQL_PASSWORD; then
    fill_val $DOCKER MYSQL_PASSWORD "$(gen_pwd)"
fi

##################
# backend/.env
##################
BACKEND="backend/.env"
if [ ! -e $BACKEND ]; then
    touch $BACKEND
fi

fill_val $BACKEND MYSQL_HOST "$(get_val $DOCKER MYSQL_HOST)"
fill_val $BACKEND MYSQL_PORT "$(get_val $DOCKER MYSQL_PORT)"
fill_val $BACKEND MYSQL_DATABASE "$(get_val $DOCKER MYSQL_DATABASE)"
fill_val $BACKEND MYSQL_USER "$(get_val $DOCKER MYSQL_USER)"
fill_val $BACKEND MYSQL_PASSWORD "$(get_val $DOCKER MYSQL_PASSWORD)"
if ! filled $BACKEND DB_DRIVER; then
    fill_val $BACKEND DB_DRIVER "sqlite"
fi
if ! filled $BACKEND APP_KEY; then
    fill_val $BACKEND APP_KEY "$(gen_app_key)"
fi
if ! filled $BACKEND APP_DEBUG; then
    fill_val $BACKEND APP_DEBUG "1"
fi
