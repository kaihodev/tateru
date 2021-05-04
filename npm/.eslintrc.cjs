module.exports = {
  extends: [
    'eslint-config-strict/default.js',
  ],
  env: { es2021: true },
  parser: 'esprima',
  parserOptions: {
    ecmaVersion: 12,
    sourceType: 'module',
  },
};
