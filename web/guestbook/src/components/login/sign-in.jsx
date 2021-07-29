import React from "react";
import loginImg from "../../assets/sign-in.svg"

export class SignIn extends React.Component{
	constructor(props) {
		super(props);
	}

	render(){
		return (
			<div className="base-container">
				<div className="header">SIGN IN TO DK-GUESTBOOK</div>
				<div className="content">
				<div className="image"> 
					<img src={loginImg}/> 
				</div>
					<div className="form">
						<div className="form-group">
							<label htmlFor="username">Username</label>
							<input type="text" name="username" placeholder="Username"/>
						</div>
						<div className="form-group">
							<label htmlFor="pasword">Password</label>
							<input type="password" name="password" placeholder="Password"/>
						</div>
					</div>
					<div className="footer">
						<button type="button" className="btn-signIn">SIGN IN</button>
					</div>
				</div>
			</div>
		);
	}
}
