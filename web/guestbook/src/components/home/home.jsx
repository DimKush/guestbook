import React, { useEffect } from 'react';
import { Redirect, Switch, useHistory, withRouter} from 'react-router-dom';
import "./home-style.scss";
import Modal from '../modal/modal';
import HomeMain from '../home-main/HomeMain';
import { cookies } from "../../App";
import  ListsTable from "../home-lists-table/lists-table.jsx";
import  ItemsTable from "../home-items-table/items-table.jsx";
import {BrowserRouter, Route, Link} from "react-router-dom";
import CreateList from '../home-lists-create/create-list';
import EditList from '../home-list-edit/edit-list';
import LoginComponent from '../login/login.jsx'

const Navigation = ({setAuthStatus, headerDescript}) => {
	const history = useHistory();

	const handleClickLogout = async() => {
		const responce = await fetch("http://localhost:8007/auth/logout", {
			headers : {"Content-type" : "application/json"},
			credentials : "include",
		});
		console.log("Logout");
		cookies.remove("jwt");
		setAuthStatus(false);

		history.push("/login");
	}

	return(
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
	);
}


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
		return(
			<LoginComponent isAuth={isAuth} setAuthStatus={setAuthStatus}/>
		);
	}


	const DateStr = new Date(Date.now()).getTime().toExponential();

	return(
	<div className="home-base-container">
		<BrowserRouter>
			{isAuth ? <Navigation setAuthStatus={setAuthStatus} headerDescript={headerDescript}/> : null}
			<div className="home-content-container">
				<div className="base-container-main">
				<Switch>
            		<Route path="/" exact component={() => <HomeMain username={username} setHeaderDescript={setHeaderDescript}/> }/>
					<Route path="/login" exact component ={()  => <LoginComponent isAuth={isAuth} setAuthStatus={setAuthStatus}/>}/>
					<Route path="/lists" exact component= {() => <ListsTable setHeaderDescript={setHeaderDescript}/>}/>
						<Route path="/lists/create" component={() => <CreateList/>} />
						<Route path="/lists/:id" component={() =><EditList/>}/>
					<Route path="/items" component={() => <ItemsTable setHeaderDescript={setHeaderDescript}/>}/>
				</Switch>
				</div>
			</div>
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={isError}/>
		</BrowserRouter>
	</div>
	);
} 