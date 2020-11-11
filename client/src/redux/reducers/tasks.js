import constants from '../../utils/constants'

const tasks = (state = [], action) => {
  switch (action.type) {
    case constants.HOME_PAGE_LOADED:
      if (action.data.message) return [...state]
      return action.data.tasks
    case constants.ADD_TASK:
      return [...state, action.task]
    case constants.DELETE_TASK:
      return [...state.filter(task => task._id !== action.id)]
    case constants.TOGGLE_TASK:
      return state.map(task =>
        task._id === action._id ? { ...task, completed: !task.completed } : task
      )
    case constants.GET_TASKS:
      return { ...state, tasks: action.json }
    case constants.UPDATE_LIST:
    default:
      return state
  }
}

export default tasks