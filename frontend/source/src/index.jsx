import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux'
import { createBrowserHistory as createHistory } from 'history'
import { Router, Route, Switch } from 'react-router-dom';
import 'whatwg-fetch';
import configureStore from './store';
import './index.css';
import { App } from './components/App';
import * as serviceWorker from './serviceWorker';

let initialState = {};

const browserHistory = createHistory();
const store = configureStore(initialState, browserHistory);

ReactDOM.render(
    <Provider store={store}>
      <Router history={browserHistory}>
        <App>
          {/* <Header/>
          <Switch>
           <Route exact path="/" component={Home} />
            <Route path="/raccoon" component={Raccoon} />
            <Route path="/fiasco" component={Fiasco} />
            <Route path="/over9k" component={Over9k} />
            <Route path="/fqueue" component={Fqueue} />
            <Route path="/goose" component={Goose} />
            <Route path="/sign" component={Sign} />
<Route component={NotFound} />
          </Switch>
          <Footer/> */}
        </App>
      </Router>
    </Provider>,
    document.getElementById('root')
  );

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
