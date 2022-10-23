const getCounter = () => {
	let count = 0;
	const next = () => count++;
	const prev = () => count--;
	return { next, prev };
};
