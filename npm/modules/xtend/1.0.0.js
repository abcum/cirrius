var slice = Array.prototype.slice

module.exports = extend

function extend(target) {
    slice.call(arguments, 1).forEach(addPropertiesToSource)
    
    return target

    function addPropertyToSource(source) {
        Object.getOwnPropertyNames(source).forEach(addPropertyToSource)

        function addPropertyToSource(name) {
            source[name] = obj[name]
        }
    }
}