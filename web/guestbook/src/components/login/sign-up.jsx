import React from "react";
import loginImg from "../../assets/sign-up.svg"

export class SignUp extends React.Component{
	constructor(props) {
		super(props);
	}

	render(){
		return (
			<div className="base-container" ref={this.props.containerRef}>
				<div className="header">SIGN UP TO DK-GUESTBOOK</div>
				<div className="content">
				<div className="image"> 
					<img src={loginImg}/> 
				</div>
					<div className="form">
						<div className="form-group">
							<label htmlFor="fullname">Full name</label>
							<input type="text" name="fullname" placeholder="Full name (not necessary)"/>
						</div>
						<div className="form-group">
							<label htmlFor="username">Username</label>
							<input type="text" name="username" placeholder="Username (required)"/>
						</div>
						<div className="form-group">
							<label htmlFor="email">Email</label>
							<input type="text" name="email" placeholder="Email (required)"/>
						</div>
						<div className="form-group">
							<label htmlFor="password">Password</label>
							<input type="password" name="password" placeholder="Password (required)"/>
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
}
