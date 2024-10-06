# Setup Rancher UI locally

**to build the container:**<br>

_docker build -t my-rancher-ui ._<br>

**to run the container:**<br>

_docker run --privileged -p 443:443 my-rancher-ui_<br>

# cypress

**REQUIREMENT TO RUN:**<br>
_To be able to run properly the cypress, it's required to set up the password of rancher UI, I did that manually, also after do the change is required to go to cypress.config.js and change the value there._<br>

**Node version:**<br>
_v14.21.3_<br>

**Install dependencies:**<br>
_npm install_<br>

**Run automation:**<br>
_npm run test:local:run_<br>
