import React from 'react';
import './App.css';
import Monthly from './Monthly';
import {
  BrowserRouter,
  Switch,
  Route,
} from "react-router-dom";

function App() {
  return (
    <div className="App">
      <div className="App-container">
        <h1 className="App-title">Accware</h1>
        <BrowserRouter>
          <Switch>
            <Route path="/:year/:month">
              <Monthly />
            </Route>
          </Switch>
        </BrowserRouter>
      </div>
    </div>
  );
}

export default App;
