import React, { useState } from 'react'
import { cookies } from "../../App";
import Modal from "../modal/modal.jsx";


async function getUsername() {
	try {
	const responce = await fetch("http://localhost:8007/auth/user", {
			  headers : { "Content-type" : "application/json",
						"Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			});

		const content = await responce.json();

		if(content.Status === "OK"){
			return content.username;
		} else {
			// TODO: Modal error
		}
	} catch(error) {

	}
}

async function DropListBack(list_id){ 
	console.log("list_id", list_id);
	const responce = await fetch(`http://localhost:8007/api/lists/${list_id}`, {
		method : "DELETE",
		headers : { "Content-type" : "application/json",
						"Authorization" :`Bearer ${cookies.get("jwt")}`},
		credentials : "include",
	});

	const content = await responce.json();

	return {Status : content.Status, Message : content.Message}
}

export default async function DeleteList(selectedRow) {
	let result = {Status : "", Message : ""};

	const currentUser = await getUsername();

	if(currentUser !== selectedRow.owner){
		result.Status = "Error";
		result.Message = `Access denied. You don't have permissions to delete this list. Owner of the record is ${selectedRow.owner}.`;

		return result;
	}

	result = await DropListBack(selectedRow.id);

	if (result.Status !== "OK") {
		result.Message=`Cannot delete record. Backend problem : ${result.Message}`;
	}

	return result;
}