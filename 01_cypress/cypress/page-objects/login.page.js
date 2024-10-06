export class LoginPage {
  constructor() {}

  getUsernameField() {
    return cy.get('[data-testid="local-login-username"]');
  }

  getPasswordField() {
    return cy.get('[data-testid="local-login-password"]');
  }

  getSubmitLoginButton() {
    return cy.get('[data-testid="login-submit"]');
  }
}
