# Imageinit

Demo

https://user-images.githubusercontent.com/29135311/117556883-9ab45280-b03b-11eb-85a2-e32969e59424.mp4

In order, here is what I did in the video:
1. Created a user "test" with password "test" and logged in
2. Added two images then logged out
3. Created another user "test2" with password "test"
4. Added two different images and logged out
5. Logged back into "test" to showcase how the previously uploaded images were unaffected
6. Deleted the two images in "test"

## Features
- Account creation
- Adding images
- Deleting images
- All images are user-specific
## Setup Locally
This guide currently only supports Linux.
### Frontend
#### AWS
From the AWS Console, obtain your Access Key, Secret Key and Cognito Identity ID. Save these in your `.env` file like this:
```
REACT_APP_AWS_ACCESS_KEY=<YOUR_ACCESS_KEY>
REACT_APP_AWS_SECRET_KEY=<YOUR_SECRET_KEY>
REACT_APP_COGNITO_IDENTITY_ID=<YOUR_COGNITO_IDENTITY_ID>
```
Then make sure to `source .env` before running the frontend server. Ensure your `.env` file is at the top level of the frontend directory.
#### Server
Make sure you have [yarn](https://yarnpkg.com/getting-started/install) installed on your system. Then run `yarn install` to install all the dependencies of the project.
Once installation has finished, run `yarn start` to run the server.
### Backend
#### Database
Follow [this guide](https://docs.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-mongodb) to install and run a MongoDB instance locally.
#### Configuration
The file `backend/config/config.yaml` contains all the settings for the Go backend. One of the configs refers to the database host and port, so make sure those fields match the information from the database instance being run locally.
#### Run Server
In the `backend/` folder,
```
go run server.go
```

