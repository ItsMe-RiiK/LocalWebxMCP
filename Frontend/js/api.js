// api.js - Handles all network requests to the Go backend

const API = {
	/**
	 * Fetches the list of uploaded files from the backend
	 * @returns {Promise<string[]>} Array of filenames
	 */
	async getFiles() {
		try {
			const response = await fetch('/api/files');
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			const data = await response.json();
			return data.files || [];
		} catch (error) {
			console.error("Error fetching files:", error);
			throw error;
		}
	},

	/**
	 * Uploads a single file to the backend
	 * @param {File} file - The file object to upload
	 * @returns {Promise<Object>} The server response
	 */
	async uploadFile(file) {
		const formData = new FormData();
		formData.append('file', file);

		try {
			const response = await fetch('/api/upload', {
				method: 'POST',
				body: formData
			});

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			return await response.json();
		} catch (error) {
			console.error("Error uploading file:", error);
			throw error;
		}
	}
};
