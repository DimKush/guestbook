import React, { Children } from 'react';
import "./modal_style.scss"

const Modal = ({active, setActive, head, msg, isError}) => {
	const handleOkClick = () => {
		setActive(false);
	}

	const handleSendAdminClick = () => {
	}

	return (
		<div className={active ? "modal active": "modal "} onClick={() => setActive(false)}>
			<div className={active ? "modal__content active" : "modal__content"} onClick={e => e.stopPropagation()}>
				{isError &&
					<div className="top-system-right-btn">
						<button>Send ticket to the administrator</button>
					</div>
				}
				<h1>{head}</h1>
				<p>{msg}</p>
				<div className="modal_footer">
					<button className="btn" onClick={handleOkClick}>OK</button>
				</div>
			</div>
			
		</div>
	);
};

export default Modal;