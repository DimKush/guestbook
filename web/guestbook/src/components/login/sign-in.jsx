import React from "react";
import "./style.scss"
import ServerStatus from "../server/status.jsx"
import Modal from "../modal/modal";
import { cookies } from "../../App";
import { Redirect } from "react-router-dom";




export default function SignIn({isAuth, setAuthStatus}){
	let usernameInput = React.createRef()
	let passwordInput = React.createRef() 

	const[modalActive, setModalActive] = React.useState(false);
	const[inputError, setErrorInput] = React.useState("");
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[isError, setIsError] = React.useState(false);
	
	console.log("isAuth in SIGN IN ", isAuth);
	if(isAuth) {
		console.log("redirect", isAuth);
		return <Redirect to="/"/>
	}

	const handleClick = function(){
		const signInObj = {
			"username" : usernameInput.current.value,
			"password" : passwordInput.current.value,
		}

		fetch("http://localhost:8007/auth/sign-in", {
			method: 'POST',
			body: JSON.stringify(signInObj),
			credentials: 'include',
			headers : {
				'Content-Type' : 'application/json'
			}
		}).then(responce => responce.json()).then(data => {
			if (data.Status === "Error") {
				setErrorInput(data.Message);
				setAuthStatus(false);
			} else if(data.Status === "OK"){
				
				setAuthStatus(true);
				cookies.set("jwt", data.token)
			}
		}).catch(error => {
			console.log("ERROR");
			setModalMsg("Server is dead.");
			setModalMsgHead("Error");
			setIsError(true);
			setModalActive(true);
		});
	} 

	const handleClickServerAlive = function(){
		const serverStatusMessage =  ServerStatus(setModalMsg, setIsError);
		console.log("serverStatusMessage", serverStatusMessage);
		setModalActive(true);
		setModalMsgHead("Server status");
	}
		return(
			<div className="base-container" >
				<div className="top-system-right-btn">
					<button onClick={handleClickServerAlive}>Server status</button>
				</div>
				<div className="header">SIGN IN</div>
				<div className="content">
				<div className="image"> 
				</div>
					<div className={!inputError ? "error" : "error active"}>{inputError}</div>
					<div className="form-login">
						<div className="form-login-group field">
							<input type="input" class="form-login-field" name="username" placeholder="Username" ref={usernameInput} required/>
							<label for="name" class="form-login-label">Username</label>
						</div>
						<div className="form-login-group field">
							<input type="password" class="form-login-field" name="password" placeholder="Password" ref={passwordInput} required/>
							<label for="password" class="form-login-label">Password</label>
						</div>
					</div>
					<div className="footer">
						<button type="button" className="btn" onClick={handleClick} >SIGN IN</button>
					</div>
				</div>
				<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={isError}/>
			</div>

		);
}
