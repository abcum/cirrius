import Ember from "ember";

export default Ember.Controller.extend({

    file: Ember.inject.controller(),
    request: Ember.inject.controller(),

    actions: {

        create: function() {

			let path = window.prompt("File path:");

			if (Ember.isEmpty(path)) return;

    		if ( path.charAt(path.length-1) === '/' ) return;

    		if ( !path.match(/^[a-zA-Z. /]+$/g) ) return;

			if ( this.get('model.files').map(i => i.get('path')).includes(path) ) return;

            this.store
                .createRecord('file', { path: path, project: this.get('model') } )
                .save()
                .then(model => {
                    this.transitionToRoute('file', model);
                })
            ;

        },

    }

});
