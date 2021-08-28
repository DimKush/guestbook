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

export default async function DeleteList(selectedRow) {
	let result = {Status : "", Message : ""};

	const currentUser = await getUsername();
	
	console.log("currentUser", currentUser);

	if(currentUser !== selectedRow.owner){
		result.Status = "Error";
		result.Message = `Access denied. You don't have permissions to delete with list. Owner of the record is ${selectedRow.owner}.`;

		return result;
	}
	result.Status="OK";
	result.Message="Record deleted";

	return result;
}