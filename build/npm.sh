#!/usr/bin/env sh

folder="$(pwd)"

versions=10

packages=()

bow=()
bow+=('async:::dist/async.min.js')
bow+=('bluebird:::js/browser/bluebird.min.js')

npm=()
npm+=('accounting:::accounting.min.js')
npm+=('chance:::dist/chance.min.js')
npm+=('deepmerge:::index.js')
npm+=('hashids:::lib/hashids.js')
npm+=('immutable:::dist/immutable.min.js')
npm+=('jsondiffpatch:::public/build/jsondiffpatch-full.min.js')
npm+=('merge:::merge.min.js')
npm+=('moment:::min/moment.min.js')
npm+=('odiff:::dist/odiff.umd.js')
npm+=('promiz:::promiz.min.js')
npm+=('q:::q.js')
npm+=('sugar:::release/sugar.min.js')
npm+=('validator:::validator.min.js')

for index in "${bow[@]}" ; do

	name="${index%%:::*}"
	path="${index##*:::}"

    npm view $name versions --json | grep '"' | cut -d '"' -f2 | tail -$versions | while read version; do

    	if [ ! -f "$folder/npm/modules/$name/$version.js" ]; then

			cd $folder
			bower install $name"#"$version
			cd "$folder/bower_components/$name"
			cp "$path" "$folder/npm/modules/$name/$version.js"
			rm -rf "$folder/bower_components/$name"

		fi

	done

done

for index in "${npm[@]}" ; do

	name="${index%%:::*}"
	path="${index##*:::}"

    npm view $name versions --json | grep '"' | cut -d '"' -f2 | tail -$versions | while read version; do

    	if [ ! -f "$folder/npm/modules/$name/$version.js" ]; then
	    	
			cd $folder
			npm install $name"@"$version
			cd "$folder/node_modules/$name"
			cp "$path" "$folder/npm/modules/$name/$version.js"
			rm -rf "$folder/node_modules/$name"

		fi

	done

done
