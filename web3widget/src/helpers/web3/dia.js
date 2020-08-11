// ERC 20 helpers for DIA
export const checkDiaBalance = async(web3,  abi, contractAddress) => {

}

export const checkDiaAllowancee = async(web3,  abi, contractAddress) => {
    
}

export const approveDiaForSpending = async(web3,  abi, contractAddress) => {
    
}


// get the details of the yield
export const getYieldDetails = async ( web3, abi, contractAddress) => {

	let yieldDetails = null;

	try {

		yieldDetails = {};
		const contract = new web3.eth.Contract(abi, contractAddress);
		yieldDetails.threeMonthPromille =  await contract.threeMonthPromille.call();
        yieldDetails.sixMonthPromille = await contract.sixMonthPromille.call();
        yieldDetails.nineMonthPromille = await contract.nineMonthPromille.call();

	}
	catch(error){
		console.log(error);
	}

	return yieldDetails;
}


export const stakeDia = async(duration, web3,  abi, contractAddress, receivingAddress, amount) => {

	let stakedDiaRes = null;

	try {

		const userAccount = (await web3.eth.getAccounts())[0];

        const contract = new web3.eth.Contract(abi, contractAddress);

        if (duration === "3m") {
            stakedDiaRes = await contract.methods.deposit3m(receivingAddress, amount).send({from: userAccount});
        }

        if (duration === "6m") {
            stakedDiaRes = await contract.methods.deposit6m(receivingAddress, amount).send({from: userAccount});
        }

        if (duration === "9m") {
            stakedDiaRes = await contract.methods.deposit9m(receivingAddress, amount).send({from: userAccount});
        }
	
	}
	catch(error){
		console.log(error);
	}

	return stakedDiaRes;

}


