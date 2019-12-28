import React from 'react';
import styles from './App.module.css';
import Monthly from './Monthly';
import Callback from './Callback';
import Login from './Login';
import {
  BrowserRouter,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";

function getThisMonth() {
  let today = new Date();
  return [today.getFullYear(), today.getMonth() + 1];
}

function App() {
  let [year, month] = getThisMonth();
  return (
    <div className={styles.app}>
      <h1 className={styles.title}>Accware</h1>
      <BrowserRouter>
        <Switch>
          <Route path="/:year/:month">
            <Monthly />
          </Route>
          <Route path="/login">
            <Login />
          </Route>
          <Route path="/callback">
            <Callback />
          </Route>
          <Route path="/">
            <Redirect to={`/${year}/${month}`} />
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
