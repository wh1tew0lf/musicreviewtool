import * as api from '../../api';

export default class authActions {

  static get actions() {
    return {
      signIn: 'SIGN_IN',
      signUp: 'SIGN_UP',
    }
  }

  static type(action) {
    const name = 'ACTION_AUTH_' + action.toUpperCase();
    return {
      request: `${name}_REQUEST`,
      success: `${name}_SUCCESS`,
      fail: `${name}_FAIL`,
    };
  }

  static request(action) {
    return {type: authActions.type(action).request};
  }

  static success(action, json) {
    return {type: authActions.type(action).success, json};
  }

  static fail(action, error) {
    return {type: authActions.type(action).fail, error};
  }

  static _makeRequest(actionType, promise) {
    return dispatch => {
      dispatch(authActions.request(actionType));

      return promise
        .then(json => {
          dispatch(authActions.success(actionType, json));
        }).catch(error => {
          dispatch(authActions.fail(actionType, error));
        });
    };
  }

  static signUp(login, password) {
    return authActions._makeRequest(authActions.actions.signUp, api.auth.signUp(login, password));
  }

  static signIn(login, password) {
    return authActions._makeRequest(authActions.actions.signIn, api.auth.signIn(login, password));
  }
}
