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

# Go api test

**REQUIREMENT TO RUN:**<br>
_To be able to run properly the automation, it's required to set up the password of rancher UI, I did that manually, also after do the change is required to add the values into line 72, on api_test.go_<br>

**Go version:**<br>
_go 1.17.13_<br>

**Run automation:**<br>
_go test -v_<br>

# Ansible to VM

**REQUIREMENT TO RUN:**<br>
_Since there is a long time, last time I used ansible, I'm not sure that script will work, also I'm NOT enable to use GPC free, so I'm not able to test, I used the documentation as reference to define the variables_<br>

_Set Up Google Cloud Credentials: Make sure you have a service account with the necessary permissions to create VMs and have the JSON key file downloaded. You must set the environment variable for the service account key; --> export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your-service-account-file.json"_<br>

**Setup ansible:**<br>
_pip install -r /path/to/requirements.txt_<br>

**Run automation:**<br>
_ansible-playbook gcp_vm_deploy.yml_<br>
