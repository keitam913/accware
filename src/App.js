import React from 'react';
import Monthly from './Monthly';
import Callback from './Callback';
import Login from './Login';
import styled from 'styled-components';
import {
  BrowserRouter,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";

const Container = styled.div`
  margin: 0 auto;
  width: 50rem;
`;

const Title = styled.h1`
  color: #666;
  font-weight: normal;
`;

function getThisMonth() {
  const today = new Date();
  return [today.getFullYear(), today.getMonth() + 1];
}

function App() {
  const [year, month] = getThisMonth();
  return (
    <Container>
      <Title>Accware</Title>
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
    </Container>
  );
}

export default App;
