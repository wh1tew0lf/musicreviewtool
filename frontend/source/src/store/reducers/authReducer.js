import {authActions} from '../actions';

export default function authReducer(state = {}, action) {
  console.log(action);
  switch (action.type) {
    case authActions.type(authActions.actions.signIn).request:
    case authActions.type(authActions.actions.signUp).request:
      return {...state, isRequested: true, requestType: action.type};
    case authActions.type(authActions.actions.signIn).fail:
    case authActions.type(authActions.actions.signUp).fail:
      return {isRequested: false, error: action.error, errorType: action.type};
    case authActions.type(authActions.actions.signIn).success:
      return {token: action.json.token, isRequested: false};
    case authActions.type(authActions.actions.signUp).success:
      return {token: '', data: action.json, isRequested: false};
    default:
  }
  return state;
}
