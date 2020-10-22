export default class Auth {
  static fetch(url, options) {
    options = options ? options : {};
    if (!options.headers) {
      options.headers = new Headers();
      options.headers.append('Content-Type', 'application/json');
    }
    return fetch(`api${url}`, options)
      .then(Auth.then);
  }

  static then(response) {
    if (response.ok) {
      return response.json();
    } else if ([400, 403, 409].includes(response.status) /* 4xx and 5xx || (400 <= response.status && response.status < 600) */) {
      return response.json()
        .then(json => {
          if (json.Message) {
            throw new Error(json.Message);
          } else {
            throw new Error(response.statusText);
          }
        })
        .catch(error => {
          throw new Error(error);
        });
    } else {
      throw new Error(response.statusText);
    }
  }

  static signUp(email, password) {
    return Auth
      .fetch('/user/new', {
      method: 'POST',
      body: JSON.stringify({
        email,
        password
      })
    });
  }

  static signIn(email, password) {
    return Auth
      .fetch('/user/login', {
      method: 'POST',
      body: JSON.stringify({
        email,
        password
      })
    });
  }

  static session(token) {
    return Auth.fetch(`/auth/${token}`);
  }

  static logOut(token) {
    return Auth
      .fetch(`/auth/${token}`, {method: 'DELETE'});
  }
}
