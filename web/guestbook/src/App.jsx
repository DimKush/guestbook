import {useEffect} from 'react';
import React from 'react';
import './App.scss';
import {BrowserRouter, Route, Switch} from "react-router-dom";
import Home from "./components/home/home.jsx";
import Cookies from 'universal-cookie';
import ListsTable from './components/home-lists-table/lists-table';
import LoginComponent from './components/login/login'
export const cookies = new Cookies();


export default function App() {
  const[isAuth, setAuthStatus] = React.useState(false);
  useEffect(() => {
    (
      async () => {
        const responce = await fetch("http://localhost:8007/auth/user", {
          headers : {"Content-type" : "application/json",
                     "Authorization" :`Bearer ${cookies.get("jwt")}`},
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

  return(
    <div className="App">
      <Home isAuth={isAuth} setAuthStatus={setAuthStatus}/>
     
    </div>
  );
}