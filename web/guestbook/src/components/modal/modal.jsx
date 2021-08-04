import React, { Children } from 'react';
import "./modal_style.scss"

const Modal = ({active, setActive, head, msg}) => {
	const handleOkClick = () => {
		setActive(false);
	}

	const handleSendAdminClick = () => {

	}

	return (
		<div className={active ? "modal active": "modal "} onClick={() => setActive(false)}>
			<div className={active ? "modal__content active" : "modal__content"} onClick={e => e.stopPropagation()}>
				<h1>{head}</h1>
				<p>{msg}</p>
				<div className="modal_footer">
					<button onClick={handleOkClick}>OK</button>
					<button onClick={handleSendAdminClick}>Send info to admin</button>
				</div>
			</div>
			
		</div>
	);
};

export default Modal;