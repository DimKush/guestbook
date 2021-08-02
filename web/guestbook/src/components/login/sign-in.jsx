import React from "react";
import loginImg from "../../assets/sign-in.svg"
import "./style.scss"

export default function SignIn(){
	let usernameInput = React.createRef()
	let passwordInput = React.createRef() 

	const handleClick = function(){
		var data = new FormData()

		const signInObj = {
			"username" : usernameInput.current.value,
			"password" : passwordInput.current.value,
		}

		//console.log(signInObj)

		data.append("sign-in", JSON.stringify(signInObj))

		fetch("http://localhost:8007/auth/sign-in", {
			method: "POST",
			body: data,
		})

	} 

		return (
			//<div className="base-container" ref={this.props.containerRef}>
			<div className="base-container" >
				<div className="header">SIGN IN TO DK-GUESTBOOK</div>
				<div className="content">
				<div className="image"> 
					<img src={loginImg}/> 
				</div>
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
			</div>
		);
}
