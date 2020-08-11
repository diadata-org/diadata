// ERC 20 helpers for DIA
export const getDiaBalance = async(web3,  abi, contractAddress) => {
	let userDiaBalance = 0;
	try {
		const userAccount = (await web3.eth.getAccounts())[0];
		const contract = new web3.eth.Contract(abi, contractAddress);
		userDiaBalance =  web3.utils.fromWei(await contract.methods.balanceOf(userAccount).call());

	}
	catch(error) {
		console.log(error);
	}

	return userDiaBalance;
}

export const getDiaAllowance = async(web3,  abi, erc20ContractAddress, yieldContractAddress) => {
	let userDiaAllowance = 0;

	try {
		const userAccount = (await web3.eth.getAccounts())[0];
		const contract = new web3.eth.Contract(abi, erc20ContractAddress);
		userDiaAllowance =  web3.utils.fromWei(await contract.methods.allowance(userAccount, yieldContractAddress).call());
	}
	catch(error) {
		console.log(error);
	}

	return userDiaAllowance;

}

export const approveDiaForStaking = async(web3,  abi, erc20ContractAddress, yieldContractAddress, amount) => {
	let diaApproved = undefined

	const userAccount = (await web3.eth.getAccounts())[0];
	const contract = new web3.eth.Contract(abi, erc20ContractAddress);
	diaApproved =  await contract.methods.approve(yieldContractAddress, web3.utils.toWei(amount)).send({ from: userAccount });
	
	return diaApproved;
}

export const stakeDia = async(duration, web3,  abi, contractAddress, receivingAddress, amount) => {

	let stakedDiaRes = null;

	const userAccount = (await web3.eth.getAccounts())[0];

	const contract = new web3.eth.Contract(abi, contractAddress);
	amount = web3.utils.toWei(amount);

	if (duration === "3m") {
		console.log(duration);
		stakedDiaRes = await contract.methods.deposit3m(receivingAddress, amount).send({from: userAccount});
	}

	if (duration === "6m") {
		console.log(duration);
		stakedDiaRes = await contract.methods.deposit6m(receivingAddress, amount).send({from: userAccount});
	}

	if (duration === "9m") {
		console.log(duration);
		stakedDiaRes = await contract.methods.deposit9m(receivingAddress, amount).send({from: userAccount});
	}

	return stakedDiaRes;

}

// get the details of the yield
export const getYieldDetails = async ( web3, abi, contractAddress) => {

	let yieldDetails = null;

	try {

		yieldDetails = {};
		const contract = new web3.eth.Contract(abi, contractAddress);
		yieldDetails.threeMonthPromille =  await contract.methods.threeMonthPromille.call().call();
        yieldDetails.sixMonthPromille = await contract.methods.sixMonthPromille.call().call();
        yieldDetails.nineMonthPromille = await contract.methods.nineMonthPromille.call().call();

	}
	catch(error){
		console.log(error);
	}

	return yieldDetails;
}

