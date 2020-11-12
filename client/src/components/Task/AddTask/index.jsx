import React, { Component, Fragment } from 'react'
import constants from '../../../utils/constants'
import { connect } from 'react-redux'
import { TextField, Button } from '@material-ui/core'
import axios from 'axios'

class AddTask extends Component {
  constructor(props) {
    super(props)

    this.state = {
      title: '',
      description: ''
    }

    this.handleChange = this.handleChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
  }

  async handleSubmit(e) {
    e.preventDefault()
    const { addTask } = this.props
    const { title, description } = this.state


    const task = {
      title,
      description,
    }

    console.log(task)

    this.setState({
      title: '',
      description: '',
    })
    return axios.post('http://localhost:8000/api/tasks/task', 
        task)
      .then(res => {
        if (res.data.error) return console.warn(res.data.error)
        addTask(res.data.taskSaved)
      })
  }

  handleChange(key, event) {
    this.setState({
      [key]: event.target.value,
    })
  }

  render() {
    const { title, description } = this.state
    return (
      <Fragment>
        <div className="wrap-taskForm">
          <div className="taskForm">
            <form onSubmit={this.handleSubmit} autoComplete="off">
              <TextField
                value={title}
                error={this.state.noTitle ? true : false}
                label="Title"
                onChange={ev => this.handleChange('title', ev)}
                name="title"
                fullWidth
                margin="normal"
              />
              <br />
              <TextField
                value={description}
                error={this.state.noDesc ? true : false}
                name="description"
                label="Description"
                onChange={ev => this.handleChange('description', ev)}
                margin="normal"
                fullWidth
                multiline
                rowsMax="4"
              />
              <Button
                style={{ marginTop: '10px' }}
                type="submit"
                color="primary"
                disabled={!title || !description}
              >
                Add task
              </Button>
            </form>
          </div>
        </div>
      </Fragment>
    )
  }
}

const mapDispatchToProps = dispatch => ({
  onChange: task => dispatch({ type: constants.ON_CHANGE, task }),
  addTask: task => dispatch({ type: constants.ADD_TASK, task }),
})

const mapStateToProps = state => ({
  tasks: state,
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AddTask)