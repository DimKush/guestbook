import React, { useEffect, useState } from 'react'
import "./edit-list-style.scss"
import { cookies } from "../../App";
import Modal from "../modal/modal.jsx";
import { Redirect, Link } from 'react-router-dom';
import { AiOutlineDoubleLeft } from 'react-icons/ai'

function ColumnCreateList ( {column, ref_current, blocked=false, defaultValue} ) {
	return(
		<div className="form-group">
		<span>{column}</span>
			
			<input class="form-field"
				defaultValue={defaultValue}
				type="text"
				id={column}
				ref={ref_current}
				disabled={blocked}
				/>
		</div>
	);
}


export default function EditList({match}) {
	const[idCheckboxBlocked, setIdCheckboxBlocked]= useState(true);
	const[ownerCheckboxBlocked, setOwnerCheckboxBlocked] = useState(true);
	const[Owners, setOwners] = useState([]);
	const[currentUser, setCurrentUser] = useState("");
	const[listObject, setListObject] = useState({});
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[modalActive, setModalActive] = React.useState(false);
	const[listOwner, setListOwner] = useState("");

	let idInput = React.createRef();
	let ownerInput = React.createRef();
	let titleInput = React.createRef();
	let descriptionInput = React.createRef();
	let auto_id_checkbox = React.createRef();
	let auto_owner_checkbox = React.createRef();

	console.log(match.params.id);
	useEffect(() => {
		(
		  async () => {
			const url = `http://localhost:8007/api/lists/${match.params.id}` 
			const responce = await fetch(url, 
			{
			  headers : {"Content-type" : "application/json",
						 "Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			  
			});
			const content = await responce.json();
			if(content.Status === "OK"){
				setListObject(content.Result);
			} else {
				// TODO : ERROR!
			}
		  }
		)();
		
		(async () => {
			const responce = await fetch("http://localhost:8007/auth/user", {
			  headers : { "Content-type" : "application/json",
						"Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			});

			const content = await responce.json();

			if(content.Status === "OK"){
				setCurrentUser(content.username);
			} else {
				// TODO: Modal error
			}
		})();
	}, []);

	const handleEditClick = () => {
		(
			async() => {
				const responce = await fetch( `http://localhost:8007/api/lists/${match.params.id}`,
				{
					method : "PUT",
					headers : { "Content-type" : "application/json",
								"Authorization" :`Bearer ${cookies.get("jwt")}`},
					credentials : "include",
					body : JSON.stringify({
						"title" : titleInput.current.value,
						"description" : descriptionInput.current.value
					}),
				});
				const status = await responce.status;
				if(status === 200){
					setModalMsg("List has been edited.");
					setModalMsgHead("OK");
					setModalActive(true);
				} else {
					const content = await responce.json();
					setModalMsg(content.Message);
					setModalMsgHead(content.Status);
					setModalActive(true);
				}
			}
		)();
	}



	return(
		<div className="list-card-main">
			<div className="row-form">
				<h1>Edit list</h1>
			</div>
			<div className="row-form">
				<div className="search-field id">
					<ColumnCreateList column={"Id"} ref_current={idInput} blocked={idCheckboxBlocked} defaultValue={listObject.id} />
				</div>
			</div>
			<div className="row-form">
				
				<div className="search-field title" >
					<ColumnCreateList column={"Title"} ref_current={titleInput} defaultValue={listObject.title}/>
				</div>
				<div className="search-field owner" >
				<div className="form-group">
					<span>Owner</span>
						<select className="form-field ownerSelect" ref={ownerInput} disabled={ownerCheckboxBlocked} onChange={e => setListOwner(e.target.value)}> 
							<option disabled selected value>{listObject.owner}</option>
								{
									Owners.map(Owner => (
										<option value={Owner}>
											{Owner}
										</option>
									))
								}
						</select>
					</div>
				</div>
			</div>
			<div className="row-form">
				<div className="search-field description">
					<ColumnCreateList column={"Description"} ref_current={descriptionInput} defaultValue={listObject.description }/>
				</div>
			</div>
			<div className="row-form">
				<Link to="/lists">
						<button className="control-but back"><AiOutlineDoubleLeft/><div className="but-tab-hight-text">Back</div></button>
				</Link>
				<button className="control-but" onClick={handleEditClick}>Edit List</button>
			</div>
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={false}/>
		</div>
	);
} 