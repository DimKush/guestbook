import React from 'react'
import { Redirect } from 'react-router-dom';
import "./home-style.scss"
import Modal from '../modal/modal';
import { cookies } from "../../App";

export default function Home({isAuth , setAuthStatus}) {
	const[modalActive, setModalActive] = React.useState(false);
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[isError, setIsError] = React.useState(false);
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
			</nav>
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={isError}/>
		</div>
	);
} 