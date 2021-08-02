import React from "react";
import loginImg from "../../assets/sign-up.svg"
import "./style.scss"

export default function SignUp(){
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
						<input type="input" class="form-login-field" name="fullname" placeholder="Full name" id='fullname' />
						<label for="name" class="form-login-label">Full name</label>
					</div>
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="username" placeholder="username" id='username' required/>
						<label for="name" class="form-login-label">Username</label>
					</div>
					<div className="form-login-group field">
						<input type="input" class="form-login-field" name="email" placeholder="email" id='email' required/>
						<label for="password" class="form-login-label">Email</label>
					</div>
					<div className="form-login-group field">
						<input type="password" class="form-login-field" name="password" placeholder="password" id='password' required/>
						<label for="password" class="form-login-label">Password</label>
					</div>
				</div>
				<div className="footer">
					<button type="button" className="btn">
						SIGN UP
					</button>
				</div>
			</div>
		</div>
	);
}
