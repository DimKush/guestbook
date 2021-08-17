import {useEffect} from 'react';
import React from 'react';
import './App.scss';
import {BrowserRouter, Route, Redirect} from "react-router-dom";
import SignIn from "./components/login/sign-in.jsx";
import SignUp from "./components/login/sign-up.jsx";
import Home from "./components/home/home.jsx";
import Cookies from 'universal-cookie';
import ListsTable from './components/home-lists-table/lists-table';

export const cookies = new Cookies();

function RightSightComponent({loggingActive, currentState, containerRef, onClick}) {
  
  return (
    <div className={loggingActive ? "right-side right" : "right-side left"} ref={containerRef} onClick={onClick}>
      <div className="inner-container">
         <div className="text">{currentState}</div>
      </div>
    </div>
  );
}


export default function App() {
  let current = React.createRef();
  const[isLoggingActive, setLoggingActive] = React.useState(true);
  const[currentState, setCurrentState] = React.useState(!isLoggingActive ? "Sign in" : "Sign up");
  const[isAuth, setAuthStatus] = React.useState(false);
  const[Username, setUsername] = React.useState("");
  
  useEffect(() => {
    (
      async () => {
        const responce = await fetch("http://localhost:8007/auth/user", {
          headers : {"Content-type" : "application/json"},
          credentials : "include",
        });
        
        const content = await responce.json();

        console.log(cookies);
        if(content.Status === "OK"){
          setAuthStatus(true);
        } else {
          setAuthStatus(false);
        }
      }
    )();

  })

  const changeState = () => {
    setLoggingActive(!isLoggingActive);
    setCurrentState(isLoggingActive ? "Sign in" : "Sign up");
  }
  
  const LoginComponent = () => {
    return(
    <main className="form-signin">
      <div className="login">
          <div className="container">
            {isLoggingActive && <SignIn containerRef={(ref) => current = ref} isAuth={isAuth} setAuthStatus={setAuthStatus} />}
            {!isLoggingActive && <SignUp containerRef={(ref) => current = ref} setLoggingActive={setLoggingActive} />}
          </div>
          <RightSightComponent loggingActive={isLoggingActive} currentState={currentState} containerRef={ref => current = ref} onClick={changeState}/>
      </div>
    </main>
    );
  }

  return (
    <div className="App">
       <BrowserRouter>          
            <Route path="/" exact component={() => <Home isAuth={isAuth} setAuthStatus={setAuthStatus}/>}/>
            <Route path="/login" component={() => <LoginComponent/>}/>
            <Route path="/lists">
              <Redirect to="/"/>
            </Route>
      </BrowserRouter>
    </div>
  );
}