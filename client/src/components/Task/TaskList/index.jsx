import React, { Component, Fragment } from 'react'
import { connect } from 'react-redux'
import axios from 'axios'
import { TaskElement } from '..'
import constants from '../../../utils/constants'

export class TaskList extends Component {
  constructor(props) {
    super(props)

    this.handleDelete = this.handleDelete.bind(this)
    this.handleToggle = this.handleToggle.bind(this)
  }

  pending(tasks) {
    const completed = tasks.filter(task => {
      if (task.completed) return task
    }).length
    return { completed, total: tasks.length }
  }

  componentDidMount() {
    const { onLoad } = this.props
    axios(`http://localhost:8000/api/tasks`)
      .then(res => {
        onLoad(res.data)
      })
      .catch(res => console.warn(`No data: ${res} `))
  }

  updateList() {
    const { onLoad } = this.props
    return axios(`http://localhost:8000/api/tasks`)
      .then(res => onLoad(res.data))
      .catch(res => console.warn(`No data: ${res} `))
  }

  handleToggle(task) {
    const { toggleTask } = this.props
    if (task.Status == "DONE") return null
    if (task.Status == "TODO") {
    axios
      .patch(`http://localhost:8000/api/tasks/task?taskId=${task.ID}`, {"status": "IN_PROGRESS"})
      .then(() => this.updateList(), toggleTask(task.ID))
    }
    if (task.Status == "IN_PROGRESS") {
      axios
      .patch(`http://localhost:8000/api/tasks/task?taskId=${task.ID}`, {"status": "DONE"})
      .then(() => this.updateList(), toggleTask(task.ID))
    }
    
  }

  handleDelete(id) {
    const { onDelete } = this.props

    return axios
      .delete(`http://localhost:8000/api/tasks/task?taskId=${id}`)
      .then(() => this.updateList(), onDelete(id))
  }

  render() {
    const tasks = this.props.tasks
    return tasks.length ? (
      <Fragment>
        {tasks &&
          tasks.map(task => {
            return (
              <TaskElement key={task.ID}
                last={true}
                {...task}
                onDelete={this.handleDelete}
                toggleTask={this.handleToggle}
              />
            )
          })}
      </Fragment>
    ) : (
      <div className="noTasks">There's no tasks!</div>
    )
  }
}

const mapStateToProps = state => ({
  tasks: state.tasks,
})

const mapDispatchToProps = dispatch => ({
  onLoad: data => dispatch({ type: constants.HOME_PAGE_LOADED, data }),
  toggleTask: task => dispatch({ type: constants.TOGGLE_TASK, task }),
  onDelete: id => dispatch({ type: constants.DELETE_TASK, id }),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TaskList)