import React from 'react';
import './App.css';
import Monthly from './Monthly';

function App() {
  return (
    <div className="App">
      <div className="container">
        <h1 className="app-title">Accware</h1>
        {Monthly()}
      </div>
    </div>
  );
}

export default App;
