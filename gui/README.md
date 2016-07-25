# app.cirrius.io

The web gui for Cirrius built using ember.js.

[![](https://img.shields.io/circleci/token/abf9e47762afcbbd936490819683ad44594f67b5/project/abcum/cirrius/master.svg?style=flat-square)](https://circleci.com/gh/abcum/cirrius) [![](https://img.shields.io/badge/ember--cli-2.6.3-orange.svg?style=flat-square)](https://github.com/abcum/cirrius) [![](https://img.shields.io/badge/license-Commercial-00bfff.svg?style=flat-square)](https://github.com/abcum/cirrius) 

#### Setup

- Install node - `brew install node`
- Install bower - `npm install -g bower`
- Install ember-cli - `npm install -g ember-cli@2.6.3`

#### Installing

- Clean cache - `npm cache clean && bower cache clean`
- Clean build - `rm -rf node_modules bower_components dist tmp`
- Install application dependencies - `npm install && bower install`

#### Upgrading

- Clean cache - `npm cache clean && bower cache clean`
- Clean build - `rm -rf node_modules bower_components dist tmp`
- Upgrade project ember-cli `npm install --save-dev ember-cli@2.6.3`
- Install application dependencies - `npm install && bower install`
- Initialise ember - `ember init`

#### Development

- Serve application - `ember serve`

#### Testing

- Initialise tests - `npm test`

#### Deployment

- Deploy production app by pushing to master branch on github.com
