/**
 * Service to abstract Wails Go backend API communication.
 */
 const wailsApiService = {
  /**
   * Reads the repository list file.
   * @returns {Promise<Array>} An array of repository settings.
   */
  async readRepositoryListFile() {
    try {
      const content = await window.go.main.App.ReadRepositoryListFile();
      return JSON.parse(content);
    } catch (error) {
      console.error("Error reading repository list file:", error);
      throw new Error(`Failed to read repository list file.`);
    }
  },

  /**
   * Writes to the repository list file.
   * @param {string} content - The JSON string to write.
   * @returns {Promise<void>}
   */
  async writeRepositoryListFile(content) {
    try {
      await window.go.main.App.WriteRepositoryListFile(content);
    } catch (error) {
      console.error("Error writing repository list file:", error);
      throw new Error(`Failed to write repository list file.`);
    }
  },

  /**
   * Uploads a single file.
   * @param {Object} repository - Repository information.
   * @param {Object} fileData - Upload file data.
   * @returns {Promise<string>} An upload result message.
   */
  async singleFileUpload(repository, fileData) {
    try {
      await window.go.main.App.AddServerXmlToFile(repository);
      await window.go.main.App.DecodeFile(fileData);
      await window.go.main.App.FileUpload("single", repository, fileData);
      await window.go.main.App.DeleteTmpFiles();
    } catch (error) {
      console.error("Error upload failed:", error);
      throw new Error(`Upload failed.`);
    }
  },

  /**
   * Create settings.xml file for upload.
   * @param {*} repository - Repository information.
   */
  async createSettingsXmlFile(repository) {
    try {
      await window.go.main.App.AddServerXmlToFile(repository);
    } catch (error) {
      console.error("Error create settings.xml file:", error);
      throw new Error(`Upload failed.`);
    }
  },

  /**
   * Delete all files of tmp direcotry.
   */
  async deleteTmpFiles() {
    try {
      await window.go.naub.App,DeleteTmpFiles();
    } catch (error) {
      console.error("Error delete tmp direcotry files:", error);
    }
  },

  /**
   * Upload multi file.
   * @param {Object} repository - Repository information.
   * @param {Object} fileData - Upload file data.
   * @returns {Promise<string>} An upload result message.
   */
  async multiFileUpload(repository, fileData) {
    try {
      await window.go.main.App.AddServerXmlToFile(repository);
      await window.go.main.App.FileUpload("multi", repository, fileData);
    } catch (error) {
      console.error("Error upload failed:", error);
      throw new Error(`Upload failed.`);
    }
  },
  /**
   * application initi
   */
  async initialize() {
    try {
      await window.go.main.App.Initialize();
    } catch (error) {
      console.error("Error initialize:", error)
      throw new Error(`Failed to initialize.`);
    }
  }
};

export default wailsApiService;
