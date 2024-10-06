const { defineConfig } = require("cypress");

module.exports = defineConfig({
  chromeWebSecurity: false,
  video: false,
  retries: 1,
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
    // Specify where your test files are located
    specPattern: "cypress/e2e/**/*.cy.{js,jsx,ts,tsx}",
    baseUrl: "https://localhost:443", // this is specific to e2e as well
    supportFile: false,
    viewportWidth: 1024,
    viewportHeight: 768,
    defaultCommandTimeout: 10000,
    env: {
      username: "admin",
      password: "1234567abcdef",
    },
  },
});
