import React from "react";
import loginImg from "../../assets/sign-in.svg"

export class SignIn extends React.Component{
	constructor(props) {
		super(props);
	}

	render(){
		return (
			<div className="base-container" ref={this.props.containerRef}>
				<div className="header">SIGN IN TO DK-GUESTBOOK</div>
				<div className="content">
				<div className="image"> 
					<img src={loginImg}/> 
				</div>
					<div className="form-login">
						<div className="form-login-group field">
							<input type="input" class="form-login-field" name="username" placeholder="Username" id='name' required/>
							<label for="name" class="form-login-label">Username</label>
						</div>
						<div className="form-login-group field">
							<input type="password" class="form-login-field" name="password" placeholder="Password" id='password' required/>
							<label for="password" class="form-login-label">Password</label>
						</div>
					</div>
					<div className="footer">
						<button type="button" className="btn">SIGN IN</button>
					</div>
				</div>
			</div>
		);
	}
}
