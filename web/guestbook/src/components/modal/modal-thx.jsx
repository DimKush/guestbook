import React, { Children } from 'react';
import "./modal_style.scss"
import {Redirect} from "react-router-dom";

const ModalThx = ({active, setActive}) => {
	const handleOkClick = () => {
		setActive(false);		
	}

	
	return(
		<div className={active ? "modal active": "modal "} onClick={() => setActive(false)}>
			<div className={active ? "modal_thx_content active" : "modal_thx_content"} onClick={e => e.stopPropagation()}>
				<div className="header-part">
					<h1>THANKS FOR REGISTERING</h1>
				</div>
				<div className="textDescr-part">
					<p>Golang/React service template text. Golang/React service template text. Golang/React service template text. Golang/React service template text.</p>
					<p>Golang/React service template text. Golang/React service template text. Golang/React service template text. Golang/React service template text.</p>
					<p>Golang/React service template text. Golang/React service template text. Golang/React service template text. Golang/React service template text.</p>
					<p>Golang/React service template text. Golang/React service template text. Golang/React service template text. Golang/React service template text.</p>
				</div>
				
				<div className="modal_footer">
					<button className="btn" onClick={handleOkClick}>OK</button>
				</div>
			</div>
			
		</div>
	);
};

export default ModalThx;