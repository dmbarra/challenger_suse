import { LoginPage } from "../page-objects/login.page";

const loginPage = new LoginPage();

before("Open Page", function () {
  cy.visit("/");
});

describe("Basic navigation on Rancher UI", () => {
  it("Success Login", () => {
    loginPage.getUsernameField().type(Cypress.env("username"));
    loginPage.getPasswordField().type(Cypress.env("password"));
    loginPage.getSubmitLoginButton().click();
    cy.url().should("contain", "dashboard/home");
    cy.title().should("eq", "Rancher");
  });
});
