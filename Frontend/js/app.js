// app.js - Handles DOM manipulation and user events

document.addEventListener('DOMContentLoaded', () => {
	// --- DOM Elements ---
	const dropZone = document.getElementById('drop-zone');
	const fileInput = document.getElementById('file-input');
	const browseBtn = document.getElementById('browse-btn');
	const fileListContainer = document.getElementById('file-list');
	const refreshBtn = document.getElementById('refresh-btn');

	// --- Initialization ---
	loadFiles();

	// --- Functions ---

	/**
	 * Fetches files from API and renders them in the grid
	 */
	function getFileIcon(fileName) {
		const ext = fileName.split('.').pop().toLowerCase();
		const imageExts = ['png', 'jpg', 'jpeg', 'gif', 'svg', 'webp'];

		if (imageExts.includes(ext)) return 'ph-image';
		if (ext === 'pdf') return 'ph-file-pdf';
		if (ext === 'json') return 'ph-file-code';
		return 'ph-file-text';
	}

	/**
	 * Fetches files from API and renders them in the grid
	 */
	async function loadFiles() {
		fileListContainer.innerHTML = '<div class="empty-state">Loading files...</div>';

		try {
			const files = await API.getFiles();

			if (files.length === 0) {
				fileListContainer.innerHTML = `
                    <div class="empty-state">
                        <i class="ph ph-folder-open" style="font-size: 2rem; margin-bottom: 0.5rem; display: block;"></i>
                        No files found. Upload something to get started!
                    </div>`;
				return;
			}

			// Clear loading state
			fileListContainer.innerHTML = '';

			// Render clickable file cards
			files.forEach(fileName => {
				const iconClass = getFileIcon(fileName);
				const fileUrl = `/api/files/${encodeURIComponent(fileName)}`;

				// Use an anchor tag <a> so it acts as a real link
				const card = document.createElement('a');
				card.className = 'file-card';
				card.href = fileUrl;
				card.target = '_blank'; // Opens in a new tab
				card.rel = 'noopener noreferrer'; // Security best practice

				card.innerHTML = `
                    <i class="ph ${iconClass}"></i>
                    <span class="file-name" title="${fileName}">${fileName}</span>
                    <i class="ph ph-arrow-square-out open-icon"></i>
                `;
				fileListContainer.appendChild(card);
			});
		} catch (error) {
			fileListContainer.innerHTML = `<div class="empty-state" style="color: var(--error)">Failed to load files. Is the backend running?</div>`;
		}
	}

	/**
	 * Handles the upload process and refreshes the UI
	 */
	async function handleFileUpload(file) {
		if (!file) return;

		// Visual feedback during upload
		const originalText = dropZone.querySelector('h3').innerText;
		dropZone.querySelector('h3').innerText = `Uploading ${file.name}...`;
		dropZone.style.opacity = '0.7';
		dropZone.style.pointerEvents = 'none';

		try {
			await API.uploadFile(file);
			// Refresh the list immediately after a successful upload
			await loadFiles();
		} catch (error) {
			alert(`Failed to upload ${file.name}. Please try again.`);
		} finally {
			// Restore UI state
			dropZone.querySelector('h3').innerText = originalText;
			dropZone.style.opacity = '1';
			dropZone.style.pointerEvents = 'auto';
		}
	}

	// --- Event Listeners ---

	// 1. Refresh Button Click
	refreshBtn.addEventListener('click', () => {
		// Add a spinning animation to the icon
		const icon = refreshBtn.querySelector('i');
		icon.style.animation = 'spin 1s linear';
		setTimeout(() => icon.style.animation = '', 1000); // Remove after 1s

		loadFiles();
	});

	// 2. Browse Button Click (Triggers the hidden file input)
	browseBtn.addEventListener('click', () => {
		fileInput.click();
	});

	// 3. File Input Change (User selected a file via dialog)
	fileInput.addEventListener('change', (e) => {
		const file = e.target.files[0];
		handleFileUpload(file);
		// Reset input so the same file can be uploaded again if needed
		fileInput.value = '';
	});

	// 4. Drag & Drop Events
	dropZone.addEventListener('dragover', (e) => {
		e.preventDefault(); // Prevent default browser behavior (opening the file)
		dropZone.classList.add('dragover');
	});

	dropZone.addEventListener('dragleave', (e) => {
		e.preventDefault();
		dropZone.classList.remove('dragover');
	});

	dropZone.addEventListener('drop', (e) => {
		e.preventDefault();
		dropZone.classList.remove('dragover');

		if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
			// Handle the first dropped file
			const file = e.dataTransfer.files[0];
			handleFileUpload(file);
		}
	});
});
