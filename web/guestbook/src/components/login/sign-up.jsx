import React from "react";
import loginImg from "../../assets/sign-up.svg"
import Modal from "../modal/modal";
import ServerStatus from "../server/status.jsx"
import "./style.scss"

export default function SignUp(){
	let fullnameInput = React.createRef();
	let usernameInput = React.createRef();
	let emailInput 	  = React.createRef();
	let passwordInput = React.createRef();
	
	//let inputError = "";

	const[modalActive, setModalActive] = React.useState(false);
	const[errorMsg, setErrorMsg] = React.useState("");
	const[inputError, setErrorInput] = React.useState("");
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[isError, setIsError] = React.useState(false);

	const handleClickServerAlive = function(){
		ServerStatus(setModalMsg, setIsError);
		console.log("Is error", isError)
		setModalActive(true);
		setModalMsgHead("Server status");
	}

	const handleClick = function(){
		const registrationDateStr = new Date(Date.now()).toISOString();
		console.log(registrationDateStr);
		
		if(!usernameInput.current.value) {
			setErrorInput("Username is required");
			return;
		} else if(!emailInput.current.value) {
			setErrorInput("Email is required");
			return;
		} else if(!passwordInput.current.value){
			setErrorInput("Password is required");
			return;
		}
		const signUpObj = {
			'name' : fullnameInput.current.value,
			'username' : usernameInput.current.value,
			'email' : emailInput.current.value,
			'password': passwordInput.current.value,
			'registration_date' : registrationDateStr,
		};

		fetch("http://localhost:8007/auth/sign-up", {
			method: 'POST',
			body: JSON.stringify(signUpObj),
			headers : {
				'Content-Type' : 'application/json'
			}
		}).then(responce => responce.json()).then(data => {
			if (data.status === "Error"){
				setErrorMsg(data.message);
				setModalMsgHead("Error");
				setModalActive(true);

			} 
		});
	}

	return (
		//<div className="base-container" ref={this.props.containerRef}>
		<div className="base-container">
			<div className="top-system-right-btn">
					<button onClick={handleClickServerAlive}>Server status</button>
			</div>
			<div className="header">SIGN UP</div>
			<div className="content">
			<div className="image"> 
				{/* <img src={loginImg}/>  */}
			</div>
				<div className={!inputError ? "error" : "error active"}>{inputError}</div>
				<div className="form-login">
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="fullname" placeholder="Full name" ref={fullnameInput} />
						<label for="name" class="form-login-label">Full name</label>
					</div>
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="username" placeholder="username" ref={usernameInput} required/>
						<label for="name" class="form-login-label">Username*</label>
					</div>
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="email" placeholder="email" ref={emailInput}  required/>
						<label for="password" class="form-login-label">Email*</label>
					</div>
					<div className="form-login-group field">
						<input type="password" class="form-login-field" name="password" placeholder="password" ref={passwordInput} required/>
						<label for="password" class="form-login-label">Password*</label>
					</div>
				</div>
				<div className="footer">
					<button type="button" className="btn" onClick={handleClick} >
						SIGN UP
					</button>
				</div>
			</div>
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={isError}/>
		</div>
	);
}
