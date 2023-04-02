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

####################################
# .env.root (single source of truth)
####################################
ENVROOT=".env.root"
if [ ! -e $ENVROOT ]; then
    cp .env.example $ENVROOT
fi

# Generate secret variable for local use
if ! filled $ENVROOT MYSQL_PASSWORD; then
    fill_val $ENVROOT MYSQL_PASSWORD "$(gen_pwd)"
fi

##################
# .env.docker
##################
ENVFILE=".env.docker"
if [ ! -e $ENVFILE ]; then
    touch $ENVFILE
fi

fill_val $ENVFILE MYSQL_DATABASE "$(get_val $ENVROOT MYSQL_DATABASE)"
fill_val $ENVFILE MYSQL_USER "$(get_val $ENVROOT MYSQL_USER)"
fill_val $ENVFILE MYSQL_PASSWORD "$(get_val $ENVROOT MYSQL_PASSWORD)"

##################
# backend/.env
##################
ENVFILE="backend/.env"
if [ ! -e $ENVFILE ]; then
    touch $ENVFILE
fi

fill_val $ENVFILE MYSQL_HOST "$(get_val $ENVROOT MYSQL_HOST)"
fill_val $ENVFILE MYSQL_PORT "$(get_val $ENVROOT MYSQL_PORT)"
fill_val $ENVFILE MYSQL_DATABASE "$(get_val $ENVROOT MYSQL_DATABASE)"
fill_val $ENVFILE MYSQL_USER "$(get_val $ENVROOT MYSQL_USER)"
fill_val $ENVFILE MYSQL_PASSWORD "$(get_val $ENVROOT MYSQL_PASSWORD)"
if ! filled $ENVFILE DB_DRIVER; then
    fill_val $ENVFILE DB_DRIVER "sqlite"
fi
if ! filled $ENVFILE APP_KEY; then
    fill_val $ENVFILE APP_KEY "$(gen_app_key)"
fi
if ! filled $ENVFILE APP_DEBUG; then
    fill_val $ENVFILE APP_DEBUG "1"
fi
