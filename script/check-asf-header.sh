#!/bin/bash
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

# List of patterns to ignore using regular expressions
IGNORED_PATTERNS=("Makefile" "NOTICE" "README.md" "README_CN.md" "SECURITY.md" "go.mod" "go.sum" "DISCLAIMER" "LICENSE" "*.json" "*.svg" "*.png" "*.jpg" "*.ico" "plugin_list" "answer")

# List of directories to ignore
IGNORED_DIRECTORIES=("answer-data" "release" "node_modules" "build")

# Recursive function to traverse directories and add ASF header
traverse_directory() {
  local dir="$1"

  # Check if the current directory needs to be ignored
  if is_directory_ignored "$dir"; then
    return
  fi

  # Iterate over all files and directories in the current directory
  for file in "$dir"/*; do
    if [ -d "$file" ]; then
      # If it's a directory, recursively process the subdirectory
      traverse_directory "$file"
    elif [ -f "$file" ]; then
      # If it's a file, check if ASF header needs to be added
      process_file "$file"
    fi
  done
}

process_file() {
  local file="$1"
  local filename=$(basename "$file")

  # Check if the file needs to be ignored
  if is_file_ignored "$filename"; then
    return
  fi

  # Check if the file already contains ASF header
  if has_asf_header "$file"; then
    echo "File $file already contains ASF header. Skipping."
    return
  fi

  # Prompt the user to add ASF header
  echo "ASF header needs to be added to file: $file"
  read -p "Do you want to add ASF header? (y/n): " choice

  if [[ $choice == "y" || $choice == "Y" ]]; then
    # Write ASF header to a temporary file
    local temp_file=$(mktemp)
    cat << EOF > "$temp_file"
/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

EOF

    # Append the original file content to the temporary file
    cat "$file" >> "$temp_file"

    # Overwrite the original file with the content of the temporary file
    mv "$temp_file" "$file"

    echo "ASF header added to file: $file"
  else
    echo "Skipping file: $file"
  fi
}

# Check if the file needs to be ignored
is_file_ignored() {
  local filename="$1"
  for pattern in "${IGNORED_PATTERNS[@]}"; do
    if [[ "$filename" == $pattern ]]; then
      return 0
    fi
  done
  return 1
}

# Check if the directory needs to be ignored
is_directory_ignored() {
  local directory="$1"
  local dirname=$(basename "$directory")
  for ignored_dir in "${IGNORED_DIRECTORIES[@]}"; do
    if [ "$dirname" = "$ignored_dir" ]; then
      return 0
    fi
  done
  return 1
}

# Check if the file already contains ASF header
has_asf_header() {
  local file="$1"
  local header=$(head -n 15 "$file")  # Assuming ASF header is within the first 15 lines
  if [[ $header == *"Licensed to the Apache Software Foundation (ASF)"* ]]; then
    return 0
  else
    return 1
  fi
}

# Execute the script, starting from the current directory
traverse_directory "../"
