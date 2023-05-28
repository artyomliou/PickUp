/**
 * @typedef {Object} Store
 * @property {number} id
 * @property {string} name
 * @property {string} openedAt
 * @property {string} closedAt
 * @property {string} pic
 * @property {string} status
 * @property {Category[]} categories
 */

/**
 * @typedef {Object} Category
 * @property {number} id
 * @property {string} name
 * @property {Product[]} products
 */

/**
 * @typedef {Object} Product
 * @property {number} id
 * @property {number} storeId
 * @property {string} name
 * @property {number} price
 * @property {string} intro
 * @property {SelectQuestion[]} [selectQuestions]
 * @property {CustomQuestion[]} [customQuestions]
 */

/**
 * @typedef {Object} SelectQuestion
 * @property {number} id
 * @property {string} name
 * @property {boolean} isRequired
 * @property {number} atMost
 * @property {SelectOption[]} options
 */

/**
 * @typedef {Object} SelectOption
 * @property {number} id
 * @property {string} name
 * @property {number} price
 * @property {number} selectQuestionId
 */

/**
 * @typedef {Object} SelectAnswer
 * @property {number} selectQuestionId
 * @property {number[]} options
 */

/**
 * @typedef {Object} CustomQuestion
 * @property {number} id
 * @property {string} name
 * @property {string} placeholder
 */

/**
 * @typedef {Object} CustomAnswer
 * @property {number} customQuestionId
 * @property {string} text
 */

/**
 * @typedef {Object} CartItem
 * @property {number} [id]
 * @property {number} [cartId]
 * @property {number} [productId]
 * @property {number} amount
 * @property {Product} [product]
 * @property {SelectAnswer[]} selectAnswers
 * @property {CustomAnswer[]} customAnswers
 */

/**
 * @typedef {Object} Order
 * @property {number} id
 * @property {number} storeId
 * @property {number} userId
 * @property {number} cartId
 * @property {number} price
 * @property {number} status
 * @property {string} createdAt
 * @property {string} updatedAt
 */
