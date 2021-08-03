import React from "react";
import loginImg from "../../assets/sign-up.svg"
import "./style.scss"

export default function SignUp(){
	let fullnameInput = React.createRef();
	let usernameInput = React.createRef();
	let emailInput 	  = React.createRef();
	let passwordInput = React.createRef();

	const handleClick = function(){
		const registrationDateStr = new Date(Date.now()).toISOString();
		console.log(registrationDateStr);

		const signUpObj = {
			'name' : fullnameInput.current.value,
			'username' : usernameInput.current.value,
			'email' : emailInput.current.value,
			'password': passwordInput.current.value,
			'registration_date' : registrationDateStr,
		};
		
		console.log(signUpObj);

		fetch("http://localhost:8007/auth/sign-up", {
			method: 'POST',
			body: JSON.stringify(signUpObj),
			headers : {
				'Content-Type' : 'application/json'
			}
		});
	}

	

	return (
		//<div className="base-container" ref={this.props.containerRef}>
		<div className="base-container">
			<div className="header">SIGN UP TO DK-GUESTBOOK</div>
			<div className="content">
			<div className="image"> 
				<img src={loginImg}/> 
			</div>
				<div className="form-login">
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="fullname" placeholder="Full name" ref={fullnameInput} />
						<label for="name" class="form-login-label">Full name</label>
					</div>
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="username" placeholder="username" ref={usernameInput} required/>
						<label for="name" class="form-login-label">Username</label>
					</div>
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="email" placeholder="email" ref={emailInput}  required/>
						<label for="password" class="form-login-label">Email</label>
					</div>
					<div className="form-login-group field">
						<input type="password" class="form-login-field" name="password" placeholder="password" ref={passwordInput} required/>
						<label for="password" class="form-login-label">Password</label>
					</div>
				</div>
				<div className="footer">
					<button type="button" className="btn" onClick={handleClick} >
						SIGN UP
					</button>
				</div>
			</div>
		</div>
	);
}
