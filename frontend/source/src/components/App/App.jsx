import React, { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import logo from '../../logo.svg';
import { authActions } from '../../store/actions';
import './App.css';

function App() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const counter = useSelector(state => state);
  const dispatch = useDispatch();

  const changeEmail = (event) => {
    setEmail(event.target.value);
  };

  const changePass = (event) => {
    setPassword(event.target.value);
  };

  const signUp = (e) => {
    e.preventDefault();
    dispatch(authActions.signUp(email, password));
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>

      <form>
        <input name="email" onChange={changeEmail} />
        <input name="email" type="password" onChange={changePass} />
        <button type="submit" onClick={signUp}>Sign Up</button>
      </form>
    </div>
  );
}

export default App;
