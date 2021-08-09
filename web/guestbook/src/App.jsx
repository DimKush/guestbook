import {useEffect} from 'react';
import React from 'react';
import './App.scss';
import {BrowserRouter, Route, Redirect} from "react-router-dom";
import SignIn from "./components/login/sign-in.jsx";
import SignUp from "./components/login/sign-up.jsx";
import Home from "./components/home/home.jsx";
import Cookies from 'universal-cookie';

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

  useEffect(() =>{
    (
      async () => {
        const responce = await fetch("http://localhost:8007/auth/user", {
          method : "GET",
          headers : {"Content-type" : "application/json"},
          credentials : "include",
        });

        const content = await responce.json();
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
    <div className="login">
        <div className="container">
          {isLoggingActive && <SignIn containerRef={(ref) => current = ref} isAuth={isAuth} setAuthStatus={setAuthStatus} />}
          {!isLoggingActive && <SignUp containerRef={(ref) => current = ref} />}
        </div>
        <RightSightComponent loggingActive={isLoggingActive} currentState={currentState} containerRef={ref => current = ref} onClick={changeState}/>
    </div>
    );
  }

  return (
    <div className="App">
       <BrowserRouter>
          {/* <Nav name={name} setName={setName}/> */}
          <main className="form-signin">
            <Route path="/" exact component={() => <Home isAuth={isAuth}/>}/>
            <Route path="/login" component={() => <LoginComponent/>}/>
          </main>
      </BrowserRouter>
    </div>
  );
}