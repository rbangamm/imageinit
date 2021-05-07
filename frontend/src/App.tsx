import React from 'react';
import './App.css';
import Dashboard from './pages/Dashboard';
import Home from './pages/Home';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import { UserContext } from './hooks/UserContext';
import useUserActions from './hooks/userActions';
import PrivateRoute from './components/PrivateRoute';


function App() {
  
  const {
    user,
    setUser,
    isLoading,
    setLoading,
    token,
    setToken
  } = useUserActions();

  return (
    <UserContext.Provider value={{ user, setUser, isLoading, setLoading, token, setToken }}>
    <Router>
      <Switch>
        <Route exact path="/">
          <Home/>
        </Route>
        <PrivateRoute path="/dashboard">
          <Dashboard/>
        </PrivateRoute>
      </Switch>
    </Router>
    </UserContext.Provider>
  );
}

export default App;
