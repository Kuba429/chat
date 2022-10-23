import Compressor from "compressorjs";

export const toBase64 = async (file: File): Promise<string> => {
	const compressedBlob = await compress(file);
	return new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.readAsDataURL(compressedBlob);
		reader.onload = () => resolve(reader.result as string);
		reader.onerror = (error) => reject(error);
	});
};
// implement generics todo
const compress = (file: File): Promise<Blob> => {
	return new Promise((resolve, reject) => {
		new Compressor(file, {
			success: (result) => {
				if (result instanceof Blob) {
					resolve(result);
				} else {
					resolve(new Blob([result]));
				}
			},
			quality: 0.7,
			maxWidth: 200,
			maxHeight: 200,
		});
	});
};
