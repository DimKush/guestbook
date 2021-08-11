import React, { useEffect } from 'react'
import { Redirect } from 'react-router-dom';
import "./home-style.scss"
import Modal from '../modal/modal';
import { cookies } from "../../App";
import  EventsTable from "../home-events-table/events-table.jsx"

export default function Home({isAuth , setAuthStatus}) {
	const[modalActive, setModalActive] = React.useState(false);
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[isError, setIsError] = React.useState(false);
	const[headerDescript, setHeaderDescript] = React.useState("Home");
	const[username, getUsername] =React.useState("Plug");

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
				<div className="username">
					<h1>Hello {username}. What we will do today?</h1>
				</div>
				<div className="bigButton">
					<div className="ButtonLogo"></div>
					<div className="ButtonText">Show all lists</div>
					<div className="bigButtonFooter"></div>
				</div>
				<div className="bigButton">
					<div className="ButtonLogo"></div>
					<div className="ButtonText">Show all events</div>
					<div className="bigButtonFooter"></div>
				</div>
				<div className="bigButton">
					<div className="ButtonLogo"></div>
					<div className="ButtonText">Show audit events</div>
					<div className="bigButtonFooter"></div>
				</div>
				
			</div>

			{/* <div className="home-content">
				<EventsTable/>
			</div>
			<div className="home-content">
				<EventsTable/>
			</div> */}
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={isError}/>
		</div>
	);
} 