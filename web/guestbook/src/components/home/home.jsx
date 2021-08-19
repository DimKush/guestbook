import React, { useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import "./home-style.scss";
import Modal from '../modal/modal';
import HomeMain from '../home-main/HomeMain';
import { cookies } from "../../App";
import  ListsTable from "../home-lists-table/lists-table.jsx";
import {BrowserRouter, Route} from "react-router-dom";


export default function Home({isAuth , setAuthStatus}) {
	const[modalActive, setModalActive] = React.useState(false);
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[isError, setIsError] = React.useState(false);
	const[headerDescript, setHeaderDescript] = React.useState("Home");
	const[username, setUsername] =React.useState("Plug");

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
			  setUsername(content.username);
			} else {
			  setAuthStatus(false);
			}
		  }
		)();
		
		
	})

	console.log("Home = ", isAuth);
	if (!isAuth) {
		console.log(isAuth);
		return <Redirect to="/login"/>
	}

	const handleClickLogout = async() => {
		const responce = await fetch("http://localhost:8007/auth/logout", {
			headers : {"Content-type" : "application/json"},
			credentials : "include",
		});
		cookies.remove("jwt");
		setAuthStatus(false);
		return <Redirect to="/login"/>

		// console.log("ERROR");
		// setModalMsg("Server is dead.");
		// setModalMsgHead("Error");
		// setModalActive(true);
		// setIsError(true); 
	}
	const DateStr = new Date(Date.now()).getTime().toExponential();

	return (
		<div className="home-base-container">
			<nav>
				<div className="nav-bg"></div>
					<div className="right-side-container">
						<ul>
							<div className="top-system-right-btn">
								<button onClick={handleClickLogout}>Logout</button>
							</div>
						</ul>
					</div>
					<div className="left-side-container">
						<ul>
							<div className="header-descript">{headerDescript}</div>
						</ul>
						<ul>
						</ul>
					</div>
			</nav>
			<div className="home-content-container">
			<BrowserRouter>
            	<Route path="/" exact component={() => <HomeMain username={username} setHeaderDescript={setHeaderDescript}/>} />
				<Route path="/lists" exact component={() => <ListsTable setHeaderDescript={setHeaderDescript}/>}/>
      		</BrowserRouter>
			
			</div>
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={isError}/>
		</div>
	);
} 