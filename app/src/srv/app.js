require('./helpers/env');

const Koa = require('koa');
const bodyParser = require('koa-bodyparser');
const cors = require('koa2-cors');
const helmet = require('koa-helmet');
const validate = require('koa-validate');

const logger = require('./helpers/get-logger')(__filename);
const router = require('./router');

const requests = require('./middlewares/requests');
const errors = require('./middlewares/errors');
const addDBToState = require('./middlewares/add-db-to-state');

const app = new Koa();
app.env = process.env.NODE_ENV;

app.use(requests());

validate(app);

app.use(errors());
app.use(cors({ origin: process.env.ALLOW_ORIGIN }));
app.use(bodyParser());
app.use(helmet());
app.use(addDBToState());

app.use(router.routes());
app.use(router.allowedMethods());

async function main() {
  app.listen(process.env.HTTP_PORT);
  logger.info('App started successfully on the port %s', process.env.HTTP_PORT);
}

main();
