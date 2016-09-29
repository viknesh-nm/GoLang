function ValidateForm() {
	var name = document.getElementById("name")
	var phn = document.getElementById("contact")
	var mail= document.getElementById("mail")
	
	if (name.value == "")
	{
		alert("Please enter your first name.");
		name.focus();
		return false;
	}
	else if (/\s/.test(name.value)) 
	{	
		alert("First Name contains space");
		name.focus();
		return false;
	}

	if (phn.value == "")
	{
		alert("Please enter your valid phone number .");
		phn.focus();
		return false;
	} 
	else if(isNaN(phn.value))
	{
		alert("enter a valid number");
		phn.focus();
		return false;
	}
	
	if (mail.value == "")
	{
		alert("Please enter your mail id.");
		mail.focus();
		return false;
	}
}