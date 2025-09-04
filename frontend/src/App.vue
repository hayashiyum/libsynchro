<template>
	<div id="app-container">
		<div class="main-layout">
			<div class="panel left-panel">
				<div class="card card-with-button">
					<h2 class="card-title">Repository</h2>
					<div class="repository-select-area">
						<div class="form-group select-group">
							<label for="repoSelectList">Registered Repositories</label>
							<select id="repoSelectList" v-model="selectedRepositoryId" class="text-input">
								<option value="" disabled>select a repository id</option>
								<option v-for="repo in repositoryList" :key="repo.repositoryId" :value="repo.repositoryId">
									{{ repo.repositoryId }}
								</option>
							</select>
						</div>
						<button class="action-button load-button" @click="loadRepositorySettings">Load</button>
						<button class="action-button clear-button" @click="clearRepositorySettings">Clear</button>
					</div>

					<div class="card-content">
						<div class="form-group">
							<div class="label-and-error">
								<label for="repositoryId">Repository ID</label>
								<span v-if="v$.repository.repositoryId.$error" class="error-message">{{ v$.repository.repositoryId.$errors[0].$message }}</span>
							</div>
							<input type="text" id="repositoryId" v-model="repository.repositoryId" class="text-input"
								@blur="v$.repository.repositoryId.$touch" />
						</div>
						<div class="form-group">
							<div class="label-and-error">
								<label for="description">Description</label>
							</div>
							<input type="text" id="description" v-model="repository.description" class="text-input" />
						</div>
						<div class="form-group">
							<div class="label-and-error">
								<label for="url">URL</label>
								<span v-if="v$.repository.repositoryUrl.$error" class="error-message">{{ v$.repository.repositoryUrl.$errors[0].$message }}</span>
							</div>
							<input type="text" id="url" v-model="repository.repositoryUrl" class="text-input" @blur="v$.repository.repositoryUrl.$touch" />
						</div>
						<div class="form-group">
							<div class="label-and-error">
								<label for="username">Username</label>
								<span v-if="v$.repository.username.$error" class="error-message">{{ v$.repository.username.$errors[0].$message }}</span>
							</div>
							<input type="text" id="username" v-model="repository.username" class="text-input"
								@blur="v$.repository.username.$touch" />
						</div>
						<div class="form-group">
							<div class="label-and-error">
								<label for="password">Password</label>
								<span v-if="v$.repository.password.$error" class="error-message">{{ v$.repository.password.$errors[0].$message }}</span>
							</div>
							<input type="password" id="password" v-model="repository.password" class="text-input" @blur="v$.repository.password.$touch" />
						</div>
						<button class="action-button" @click="saveRepositoryList">Save Repository</button>
					</div>
				</div>
			</div>

			<div class="panel right-panel">
				<div class="card card-with-button">
					<h2 class="card-title">Jar/Pom File</h2>
					<div class="tabs-header">
						<button :class="{ 'tab-button': true, 'active': currentTab === 'single' }" @click="currentTab = 'single'">
							Single Jar/Pom File
						</button>
						<button :class="{ 'tab-button': true, 'active': currentTab === 'multi' }" @click="currentTab = 'multi'">
							Multi Jar/Pom File
						</button>
					</div>

					<div class="card-content tab-content-wrapper">
						<div v-if="currentTab === 'single'" class="tab-pane">
							<div class="form-group">
								<div class="label-and-error">
									<label for="groupId">Group ID</label>
									<span v-if="v$.singleFileForm.groupId.$error" class="error-message">{{ v$.singleFileForm.groupId.$errors[0].$message }}</span>
								</div>
								<input type="text" id="groupId" v-model="singleFileForm.groupId" class="text-input" @blur="v$.singleFileForm.groupId.$touch"/>
							</div>
							<div class="form-group">
								<div class="label-and-error">
									<label for="artifactId">Artifact ID</label>
									<span v-if="v$.singleFileForm.artifactId.$error" class="error-message">{{ v$.singleFileForm.artifactId.$errors[0].$message }}</span>
								</div>
								<input type="text" id="artifactId" v-model="singleFileForm.artifactId" class="text-input" @blur="v$.singleFileForm.artifactId.$touch" />
							</div>
							<div class="form-group">
								<div class="label-and-error">
									<label for="version">Version</label>
									<span v-if="v$.singleFileForm.version.$error" class="error-message">{{ v$.singleFileForm.version.$errors[0].$message }}</span>
								</div>
								<input type="text" id="version" v-model="singleFileForm.version" class="text-input" @blur="v$.singleFileForm.version.$touch"/>
							</div>
							<div class="form-group">
								<div class="label-and-error">
									<label for="jarPomFile">File Path</label>
									<span v-if="v$.singleFileForm.jarPomFile.$error" class="error-message">{{ v$.singleFileForm.jarPomFile.$errors[0].$message }}</span>
								</div>
								<input type="file" id="jarPomFile" @change="handleSingleFileChange" class="file-input" />
							</div>
							<button class="action-button" @click="uploadSingleJar">Upload</button>
						</div>

						<div v-if="currentTab === 'multi'" class="tab-pane">
							<div class="form-group">
								<div class="label-and-error">
									<label for="jarPomListFile">List File Path</label>
									<span v-if="v$.multiFileForm.jarPomListFile.$error" class="error-message">{{ v$.multiFileForm.jarPomListFile.$errors[0].$message }}</span>
								</div>
								<input type="file" id="jarPomListFile" @change="handleMultiFileChange" class="file-input" />
							</div>

							<div class="multi-file-content-wrapper">
								<label>File Content</label>
								<div class="multi-file-content">
									<pre class="log-text">{{ multiFileForm.multiFileContent || 'File content will appear here after loading...' }}
									</pre>
								</div>
							</div>

							<div class="multi-button-container">
								<button class="action-button" @click="uploadMultiJar">Upload</button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<LogArea :logs='logs'></LogArea>
		<Footer></Footer>
	</div>
</template>

<script>
import { reactive, computed } from 'vue';
import useVuelidate from '@vuelidate/core';
import { required, helpers, url } from '@vuelidate/validators';
import wailsApiService from './wailsApiService.js';
import Footer from './components/Footer.vue';
import LogArea from './components/LogArea.vue';

const repositoryIdRule = helpers.regex(/^[a-zA-Z0-9_-]+$/)
const groupIdRule = helpers.regex(/^[a-z0-9_\.]+$/)
const artifactIdRule = helpers.regex(/^[a-z0-9-]+$/)

export default {
	name: 'App',
	components: {
		LogArea,
		Footer
	},
	setup() {
		// Defines reactive data models for the form inputs.
		const repository = reactive({
			repositoryId: '',
			description: '',
			repositoryUrl: '',
			username: '',
			password: '',
		});
		const singleFileForm = reactive({
			groupId: '',
			artifactId: '',
			version: '',
			jarPomFile: null,
		});
		const multiFileForm = reactive({
			jarPomListFile: null,
			multiFileContent: '',
		});

		// Defines validation rules using Vuelidate helpers.
		const rules = computed(() => {
			return {

				// repository validation rules
				repository: {
					repositoryId: {
						required: helpers.withMessage('* Repository ID is required', required),
						repositoryIdRule: helpers.withMessage('* Only alphanumeric characters, digits and symbols ("-","_") are allowed.', repositoryIdRule)
					},
					repositoryUrl: { 
						required: helpers.withMessage('* URL is required', required)
					},
					username: { required: helpers.withMessage('* Username is required', required) },
					password: { required: helpers.withMessage('* Password is required', required) },
				},
				// singleFileForm validation rules
				singleFileForm: {
					groupId: { 
						required: helpers.withMessage('* Group ID is required', required),
						groupIdRule: helpers.withMessage('* Only lowercase letters, digits, dot and underbar', groupIdRule),
					},
					artifactId: { 
						required: helpers.withMessage('* Artifact ID is required', required),
						artifactIdRule: helpers.withMessage("* Only lowercase letters, digits, and hyphens", artifactIdRule)
					},
					version: { required: helpers.withMessage('* Version is required', required) },
					jarPomFile: { required: helpers.withMessage('* File is required', required) },
				},
				// multiFileForm validation rules
				multiFileForm: {
					jarPomListFile: { required: helpers.withMessage('* File is required', required) },
				}
			};
		});
		const v$ = useVuelidate(rules, { repository, singleFileForm, multiFileForm });

		return {
			repository,
			singleFileForm,
			multiFileForm,
			v$,
		};
	},
	data() {
		return {
			currentTab: 'single',
			repositoryList: [],
			selectedRepositoryId: '',
			logs: [],
		};
	},
	async mounted() {
		// Initializes the Wails API service.
		await wailsApiService.initialize();
		// Loads the list of registered repositories.
		this.loadRepositoryList();
		// Listens for log messages from the backend and adds them to the logs array.
		window.runtime.EventsOn("logMessage", (data) => {
			if (data && data.type && data.message) {
				this.logs.push(data);
				// Automatically scrolls the log area to the bottom to show new messages.
				this.$nextTick(() => {
					const logOutput = this.$el.querySelector('.log-output');
					if (logOutput) {
						logOutput.scrollTop = logOutput.scrollHeight;
					}
				});
			}
		});
	},
	beforeUnmount() {
		// Removes the event listener before the component is unmounted.
		window.runtime.EventsOff("logMessage");
	},
	methods: {
		/**
		 * Handles the change event for the single file input, setting the selected file.
		 * @param {Event} event The file input change event.
		 */
		handleSingleFileChange(event) {
			const file = event.target.files[0];
			if (file) {
				this.singleFileForm.jarPomFile = file;
			} else {
				this.singleFileForm.jarPomFile = null;
			}
		},
		/**
		 * Handles the change event for the multi-file list input, setting the selected file and loading its content.
		 * @param {Event} event The file input change event.
		 */
		handleMultiFileChange(event) {
			const file = event.target.files[0];
			if (file) {
				this.multiFileForm.jarPomListFile = file;
				this.loadFileContent();
			} else {
				this.multiFileForm.jarPomListFile = null;
				this.multiFileForm.multiFileContent = '';
			}
		},

		/**
		 * Reads the content of the selected multi-file list and updates the UI.
		 */
		loadFileContent() {
			if (!this.multiFileForm.jarPomListFile) {
				this.multiFileForm.multiFileContent = 'No file selected to load.';
				return;
			}
			const reader = new FileReader();
			reader.onload = (e) => {
				this.multiFileForm.multiFileContent = e.target.result;
			};
			reader.onerror = () => {
				this.multiFileForm.multiFileContent = 'Error reading file.';
			};
			reader.readAsText(this.multiFileForm.jarPomListFile);
		},

		/**
		 * Loads the settings for the selected repository into the form fields.
		 */
		loadRepositorySettings() {
			// Clear error messages by resetting validation state
			this.v$.repository.$reset();

			if (!this.selectedRepositoryId) {
				this.logs.push({ type: 'INFO', message: `[INFO] Please select a repository to load` });
				return;
			}
			const selectedRepo = this.repositoryList.find(
				(repo) => repo.repositoryId === this.selectedRepositoryId
			);
			if (selectedRepo) {
				this.repository.repositoryId = selectedRepo.repositoryId;
				this.repository.repositoryUrl = selectedRepo.repositoryUrl;
				this.repository.username = selectedRepo.username || '';
				this.repository.password = selectedRepo.password || '';
				this.repository.description = selectedRepo.description || '';
				this.logs.push({ type: 'INFO', message: `[INFO] Repository '${selectedRepo.repositoryId}' loaded` });
			} else {
				this.logs.push({ type: 'ERROR', message: `[ERROR] Selected repository not found` });
			}
		},


		/**
		 * Asynchronously reads the repository list file from the backend and updates the UI.
		 */
		async loadRepositoryList() {
			try {
				this.repositoryList = await wailsApiService.readRepositoryListFile();
				this.logs.push({ type: 'INFO', message: `[INFO] Repository list loaded` });
			} catch (error) {
				this.logs.push({ type: 'ERROR', message: `[ERROR] ${error.message}` });
			}
		},

		/**
		 * Validates and saves the current repository settings to a JSON file.
		 * It handles both adding new repositories and updating existing ones.
		 */
		async saveRepositoryList() {
			// Execute validation
			const isRepoValid = await this.v$.repository.$validate();
			if (!isRepoValid) {
				this.logs.push({ type: 'ERROR', message: `[ERROR] Validation error, check your input` });
				return;
			}

			const newRepo = {
				repositoryId: this.repository.repositoryId,
				repositoryUrl: this.repository.repositoryUrl,
				username: this.repository.username,
				password: this.repository.password,
				description: this.repository.description,
			};

			try {
				let parsedList = [];
				try {
					parsedList = await wailsApiService.readRepositoryListFile();
				} catch (readError) {
					this.logs.push({ type: 'ERROR', message: readError.message });
					throw readError;
				}

				const repoIndex = parsedList.findIndex(
					(repo) => repo.repositoryId === newRepo.repositoryId
				);

				if (repoIndex > -1) {
					parsedList[repoIndex] = newRepo;
					this.logs.push({ type: 'INFO', message: `[INFO] Repository '${newRepo.repositoryId}' updated` });
				} else {
					parsedList.push(newRepo);
					this.logs.push({ type: 'INFO', message: `[INFO] Repository '${newRepo.repositoryId}' added` });
				}

				const updatedContent = JSON.stringify(parsedList, null, 2);

				try {
					await wailsApiService.writeRepositoryListFile(updatedContent);
				} catch (error) {
					this.logs.push({ type: 'ERROR', message: `[ERROR] ${error.message}` });
				}
				await this.loadRepositoryList();
			} catch (error) {
				this.logs.push({ type: 'ERROR', message: `[ERROR] An error occurred during repository registration: ${error.message}` });
			}
		},

		/**
		 * Clears all repository-related form fields and resets the validation state.
		 */
		clearRepositorySettings() {
			this.repository.repositoryId = '';
			this.repository.repositoryUrl = '';
			this.repository.username = '';
			this.repository.password = '';
			this.repository.description = '';
			this.selectedRepositoryId = '';
			this.v$.repository.$reset();
		},

		/**
		 * Validates the single file upload form and initiates the upload process to the backend.
		 */
		async uploadSingleJar() {

			// Clear validation error messages
			this.v$.repository.$reset();
			this.v$.singleFileForm.$reset();

			// Validation
			const isRepoValid = await this.v$.repository.$validate();
			const isSingleFileValid = await this.v$.singleFileForm.$validate();
			if (!isRepoValid || !isSingleFileValid) {
				this.logs.push({ type: 'ERROR', message: `[ERROR] Validation error, check your input` });
				return;
			}

			this.logs.push({ type: 'INFO', message: `[INFO] Starting file upload...` });

			const repository = {
				repositoryId: this.repository.repositoryId,
				repositoryUrl: this.repository.repositoryUrl,
				username: this.repository.username,
				password: this.repository.password,
			};

      		// create file for tmp directory 
			const base64FileData = this.singleFileForm.jarPomFile
				? await this.fileToBase64(this.singleFileForm.jarPomFile)
				: null;

			const uploadFileData = {
				artifactId: this.singleFileForm.artifactId,
				groupId: this.singleFileForm.groupId,
				version: this.singleFileForm.version,
				uploadFileName: this.singleFileForm.jarPomFile.name,
				base64FileData: base64FileData,
			};


			// uploadFileData validation
			// TODO

      		// upload file
			try {
				await wailsApiService.singleFileUpload(repository, uploadFileData);
			} catch (error) {
				this.logs.push({ type: 'ERROR', message: `[ERROR] ${error.message}` });
			}
			this.logs.push({ type: 'INFO', message: `[INFO] Single file upload process completed` });
		},

		/**
		 * Converts a File object to a Base64 encoded string.
		 * @param {File} file The file to convert.
		 * @returns {Promise<string>} A promise that resolves with the Base64 string.
		 */
		fileToBase64(file) {
			return new Promise((resolve, reject) => {
				const reader = new FileReader();
				reader.onload = () => {
					const base64String = reader.result.split(',')[1];
					resolve(base64String);
				};
				reader.onerror = (error) => {
					reject(error);
				};
				reader.readAsDataURL(file);
			});
		},

		/**
		 * Validates the multi-file upload form, reads the list file, and initiates
		 * the upload process for each file listed in the file.
		 */
		async uploadMultiJar() {

			// Clear validation error messages
			this.v$.repository.$reset();
			this.v$.multiFileForm.$reset();

			// Validation
			const isRepoValid = await this.v$.repository.$validate();
			const isMultiFileValid = await this.v$.multiFileForm.$validate();
			if (!isRepoValid || !isMultiFileValid) {
				this.logs.push({ type: 'ERROR', message: `[ERROR] Validation error, check your input` });
				return;
			}

			this.logs.push({ type: 'INFO', message: `[INFO] Starting file upload...` });

			const repository = {
				repositoryId: this.repository.repositoryId,
				repositoryUrl: this.repository.repositoryUrl,
				username: this.repository.username,
				password: this.repository.password,
			};
			const uploadFileContents = this.multiFileForm.multiFileContent.split('\n');

      		// upload for each file
			for (const line of uploadFileContents) {
				const trimmedLine = line.trim();
				if (!trimmedLine) {
					continue;
				}
				const parts = trimmedLine.split(',').map(part => part.trim());
				if (parts.length < 4) {
					this.logs.push({ type: 'ERROR', message: `[ERROR] Invalid line format: ${trimmedLine}` });
					continue;
				}
				const [groupId, artifactId, version, filePath] = parts;

				try {
					this.logs.push({ type: 'INFO', message: `[INFO] Starting upload for ${groupId}:${artifactId}:${version}:${filePath}` });
					const uploadFileData = {
						groupId: groupId,
						artifactId: artifactId,
						version: version,
						uploadFilePath: filePath,
					};
					await wailsApiService.createSettingsXmlFile(repository);
					await wailsApiService.multiFileUpload(repository, uploadFileData);
					await wailsApiService.deleteTmpFiles();
				} catch (error) {
					this.logs.push({ type: 'ERROR', message: `[ERROR] Failed to upload ${groupId}:${artifactId}:${version}:${filePath} ${error.message}` });
				}
				this.logs.push({ type: 'INFO', message: `[INFO] ======================== upload next file ========================` });
			}
			this.logs.push({ type: 'INFO', message: `[INFO] Multi file upload process completed` });
		},
	},
};
</script>

<style>
/* Main layout for left and right panels */
.main-layout {
	display: flex;
	flex-wrap: wrap;
	gap: 16px;
	flex-grow: 1;
	padding: 16px;
}

/* Panel styling */
.panel {
	flex: 1;
	min-width: 300px;
	padding: 16px;
	box-sizing: border-box;
}

/* Responsive columns for md breakpoint */
@media (min-width: 768px) {

	.left-panel,
	.right-panel {
		flex: 0 0 calc(50% - 8px);
		max-width: calc(50% - 8px);
	}
}

/* Card styling with button at the bottom */
.card-with-button {
	background-color: #ffffff;
	border-radius: 8px;
	box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
	padding: 20px;
	height: 100%;
	box-sizing: border-box;
	display: flex;
	flex-direction: column;
}

.card-title {
	font-size: 1.5rem;
	margin-bottom: 20px;
	color: #333;
}

/* Updated card-content to occupy available space */
.card-content {
	padding-top: 10px;
	flex-grow: 1;
	display: flex;
	flex-direction: column;
}

/* Form group styling */
.form-group {
	margin-bottom: 15px;
}

/* New style to align labels and error messages horizontally */
.label-and-error {
	display: flex;
	align-items: center;
	gap: 10px;
	margin-bottom: 5px;
}

/* Base style for all labels */
.form-group label {
	display: block;
	margin-bottom: 5px;
	font-weight: bold;
	color: #555;
	text-align: left;
	flex-shrink: 0;
}

/* Remove bottom margin for labels inside .label-and-error */
.label-and-error label {
	margin-bottom: 0;
}

.text-input,
.file-input {
	width: 100%;
	padding: 10px;
	border: 1px solid #ccc;
	border-radius: 4px;
	box-sizing: border-box;
	font-size: 1rem;
}

.text-input:focus,
.file-input:focus {
	border-color: #007bff;
	outline: none;
	box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}

/* File input specific styling */
.file-input {
	cursor: pointer;
}

.file-name {
	display: block;
	margin-top: 5px;
	font-size: 0.9rem;
	color: #666;
}

/* Tabs styling */
.tabs-header {
	display: flex;
	margin-bottom: 20px;
	border-bottom: 1px solid #eee;
}

.tab-button {
	flex-grow: 1;
	padding: 10px 15px;
	border: none;
	background-color: transparent;
	cursor: pointer;
	font-size: 1rem;
	color: #555;
	border-bottom: 3px solid transparent;
	transition: all 0.3s ease;
}

.tab-button:hover {
	color: #007bff;
}

.tab-button.active {
	color: #007bff;
	border-bottom-color: #007bff;
	font-weight: bold;
}

.tab-pane {
	padding-top: 15px;
	flex-grow: 1;
	display: flex;
	flex-direction: column;
}

/* New style for the file content display wrapper in multi tab */
.multi-file-content-wrapper {
	margin-top: 15px;
	margin-bottom: 15px;
	flex-grow: 1;
	display: flex;
	flex-direction: column;
}

.multi-file-content-wrapper label {
	display: block;
	margin-bottom: 5px;
	font-weight: bold;
	color: #555;
	text-align: left;
}

/* Reusing log-output styles for multi-file-content */
.multi-file-content {
	min-height: 220px;
	max-height: 220px;
	overflow-y: auto;
	background-color: #cdcdcd;
	border: 1px solid #ddd;
	border-radius: 4px;
	padding: 10px;
	font-family: 'Courier New', Courier, monospace;
	white-space: pre-wrap;
	flex-grow: 1;
}

/* Action buttons styling */
.action-button {
	display: block;
	width: 100%;
	padding: 12px;
	margin-top: auto;
	font-size: 1rem;
	font-weight: bold;
	color: #fff;
	background-color: #007bff;
	border: none;
	border-radius: 4px;
	cursor: pointer;
	transition: background-color 0.3s ease;
}

.action-button:hover {
	background-color: #0056b3;
}

/* Multi-button container for flex layout */
.multi-button-container {
	display: flex;
	gap: 10px;
	margin-top: 20px;
	padding-top: 0;
}

.multi-button-container .action-button {
	margin-top: 0;
	flex-grow: 1;
}

/* New styles for repository select area */
.repository-select-area {
	display: flex;
	align-items: flex-end;
	gap: 10px;
	margin-bottom: 20px;
}

.repository-select-area .select-group {
	flex-grow: 1;
	margin-bottom: 0;
}

/* Style for the "Registered Repository" label */
.repository-select-area .select-group label {
	margin-bottom: 5px;
}

.repository-select-area .load-button {
	width: auto;
	padding: 12px 20px;
	margin-top: 0;
	flex-shrink: 0;
}

/* Style for the Clear button */
.repository-select-area .clear-button {
	width: auto;
	padding: 12px 20px;
	margin-top: 0;
	flex-shrink: 0;
	background-color: #5a6268;
}

.repository-select-area .clear-button:hover {
	background-color: #599551;
}

/* Error message styling */
.error-message {
	color: red;
	font-size: 0.75rem;
	padding-bottom: 1px;
	flex-shrink: 0;
	white-space: nowrap;
	text-align: left;
}
</style>