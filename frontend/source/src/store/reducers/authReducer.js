import {authActions} from '../actions';

export default function authReducer(state = {}, action) {
  switch (action.type) {
    case authActions.type(authActions.actions.signIn).request:
    case authActions.type(authActions.actions.signUp).request:
    case authActions.type(authActions.actions.logOut).request:
    case authActions.type(authActions.actions.check).request:
      return {...state, isRequested: true, requestType: action.type};
    case authActions.type(authActions.actions.signIn).fail:
    case authActions.type(authActions.actions.signUp).fail:
    case authActions.type(authActions.actions.logOut).fail:
    case authActions.type(authActions.actions.check).fail:
      return {isRequested: false, error: action.error, errorType: action.type};
    case authActions.type(authActions.actions.signIn).success:
      return {token: action.json.token, isRequested: false};
    case authActions.type(authActions.actions.signUp).success:
      return {token: '', data: action.json, isRequested: false};
    case authActions.type(authActions.actions.logOut).success:
      return {token: '', isRequested: false};
    case authActions.type(authActions.actions.check).success:
      return {...state, isRequested: false};
    default:
  }
  return state;
}
