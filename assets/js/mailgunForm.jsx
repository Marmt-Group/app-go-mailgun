

class MailgunForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            fileName: '',
            response: '',
            showResponse: false,
            showError: false
        }
        this.fileInput = React.createRef();
    }

    handleFileUpdate = (event) => {
        this.setState({ fileName: event.target.files[0].name });
    }

    handleDismiss = () => {
        this.setState({ showResponse: false })
    }

    handleSubmit = (event) => {
        event.preventDefault()

        const formData = new FormData()
        formData.append('file', this.fileInput.current.files[0])

       return fetch('/upload', {
            method: 'POST',
            mode: 'same-origin',
            cache: 'no-cache',
            credentials: 'same-origin',
            body: formData
        })
        .then(response => response.text())
        .then(message => {
            this.setState({ response: message, showResponse: true})
        })
        .catch(error => {
            this.setState({ response: error, showResponse: true, showError: true})
        });
    }

    render() {

        return (
            <div className="container form-group">
                <form onSubmit={this.handleSubmit}>
                    <div className="field is-grouped">
                        <div className="control">
                            <div className="file has-name">
                                <label className="file-label">
                                    <input className="file-input" type="file" onChange={this.handleFileUpdate} ref={this.fileInput} required />
                                        <span className="file-cta">
                                            <span className="file-icon">
                                                <i className="fas fa-upload"></i>
                                            </span>
                                            <span className="file-label">
                                                Choose a CSV file
                                            </span>
                                        </span>
                                        <span className="file-name">
                                            {this.state.fileName ? this.state.fileName : 'Name of file'}
                                        </span>
                                </label>
                            </div>
                        </div>

                        <div className="control">
                            <button className="button is-primary" type="Submit">Submit</button>
                        </div>
                    </div>
                </form>

                {this.state.showResponse &&
                    <div className={`notification ${this.state.showError ? 'is-danger' : 'is-primary'} `}>
                        <button className="delete" onClick={this.handleDismiss}></button>
                        {this.state.response}
                    </div>
                }
            </div>
        );
    }
}

ReactDOM.render(<MailgunForm />, document.getElementById("mailgun-form"));